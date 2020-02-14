# コンドルセのアルゴリズム
import random

c_alpha = 1
c_beta = 2
c_gamma = 1
c_delta = 1

# 返り値のテーブルはハッシュのハッシュの構造をもつ
# table[i][j] == 1　なら、validators[i] は validators[j]　より強い
# table[i][j] == -1　なら、validators[i] は validators[j]　より弱い
# table[i][j] == 0　なら、validators[i] は validators[j]　と等しい
# iは主体、jは比較対象
def create_attr_table(validators, attr):
    validators_size = len(validators)
    attr_table = [[0 for i in range(validators_size)] for j in range(validators_size)]
    for i in range(len(validators)):
        for j in range(len(validators)):
            validator = validators[i]
            target_validator = validators[j]
            if i == j or getattr(validator, attr) == getattr(target_validator, attr):
                attr_table[i][j] = 0
            elif getattr(validator, attr) > getattr(target_validator, attr):
                attr_table[i][j] = 1
            else:
                attr_table[i][j] = -1
    return attr_table

def create_total_table(validators):
    attr_table = {}
    attrs = ["staking_token", "age", "reputation", "random"]
    for attr in attrs:
        attr_table[attr] = create_attr_table(validators, attr)

    validators_size = len(validators)
    total_table = [[0 for i in range(validators_size)] for j in range(validators_size)]

    for i in range(len(validators)):
        for j in range(len(validators)):
            if i == j:
                total_table[i][j] = 0
            else:
                weights = [c_alpha, c_beta, c_gamma, c_delta]
                scores = [attr_table[attr][i][j] for attr in attrs]
                total_table[i][j] = sum(list(map(lambda x,y: x*y, weights, scores)))
    return total_table

# ノードがある場合は重み、ノードがない場合は0
# ex: graph[1][2] = 3 の場合、ノード1からノード2へ重み3のエッジがある
def create_graph(validators):
    total_table = create_total_table(validators)
    graph = [[-1]*len(total_table) for i in range(len(total_table))]
    for i in range(len(validators)):
        for j in range(len(validators)):
            score = total_table[i][j]
            if score < 0:
                graph[i][j] = -score
            else:
                graph[i][j] = 0
    return graph

# 勝者を決定する
# データを重み付き有向グラフで保持する
def winner_by_graph(graph, validators):
    sink_count = 0
    sink_validators = []
    for i in range(len(validators)):
        if len(set(graph[i])) == 1:
            sink_count += 1
            sink_validators.append(validators[i])

    # シンク(どこへの辺もない頂点)が2つ以上ある
    if sink_count > 1:
        # print("MORE THAN TWO SINK")
        # Check if one candidate has the highest score in R, if so, this candidate is the winner.
        # Reputationのスコアがもっとも高い者を勝者にする(?)
        validator_to_reputation = {v: v.reputation for v in sink_validators}
        highest_reputation_validator_list = [kv[0] for kv in validator_to_reputation.items() if kv[1] == max(validator_to_reputation.values())]
        if len(highest_reputation_validator_list) == 1: return highest_reputation_validator_list[0]

        # If there is still a tie, check the same for A, E, and U respectively. If there is only one candidate that has a higher score, there are declared the winner.
        # A,E,Uの中でもっとも多く制した者が勝ち
        attrs = ["staking_token", "age", "random"]
        attr_to_validator_to_value = {}
        for attr in attrs:
            attr_to_validator_to_value[attr] = {v: getattr(v, attr) for v in sink_validators}
        highest_validator_list = []
        for attr in attrs:
            highest_validator_list.extend([kv[0] for kv in attr_to_validator_to_value[attr].items() if kv[1] == max(attr_to_validator_to_value[attr].values())])
        validator_to_highest_count = {v: 0 for v in sink_validators}

        for validator in highest_validator_list:
            validator_to_highest_count[validator] += 1

        highest_validator_list = [kv[0] for kv in validator_to_highest_count.items() if kv[1] == max(validator_to_highest_count.values())]
        if len(highest_validator_list) == 1: return highest_validator_list[0]

        # TODO: 決まらなかったらランダムできめるので修正する(ダメそう)
        # 上の対処法としてこちらまで処理が来ないようにする。すなわち、reputationの被りがないようにする
        return random.choice(highest_validator_list)

        # TODOの代わりに、候補者のみで再度グラフを構築し再帰する
        # graph = create_graph(highest_validator_list)
        # return winner_by_graph(graph, highest_validator_list)


    # シンクが1つある
    if sink_count == 1:
        # print("ONE SINK")
        return sink_validators[0]

    # シンクがない
    # 弱い頂点を削っていく
    if sink_count == 0:
        # print("NO SINK")
        min_edge_weight = 1000000000000000
        for i in range(len(validators)):
            for j in range(len(validators)):
                if graph[i][j] < min_edge_weight and graph[i][j] > 0:
                    min_edge_weight = graph[i][j]
        for i in range(len(validators)):
            for j in range(len(validators)):
                if graph[i][j] == min_edge_weight:
                    graph[i][j] = 0

        sink_count = 0
        sink_validators = []
        for i in range(len(validators)):
            if len(set(graph[i])) == 1:
                sink_count += 1
                sink_validators.append(validators[i])

        # シンクがそれでもないときはランダムに返す、不適切そう
        if (len(sink_validators) == 0):
            return random.choice(validators)
        # 無限ループ対策、不適切そう
        if (len(sink_validators) == len(validators)):
            return random.choice(sink_validators)
        graph = create_graph(sink_validators)
        return winner_by_graph(graph, sink_validators)
