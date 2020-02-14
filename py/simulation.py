# 合意形成アリゴリズムの評価軸
# ・正しいブロックを作れるか
# ・中央集権度合い
# ・攻撃耐性 <- 今回は考えない
# ・スループット <- 今回は考えない
# ・電力消費量 <- 今回は考えない

# 手法
# ・コンドルセ
# ・ヤング
# ・多数決
# ・m票先取
# ・重み付き投票
# ・ランダム(= 1票先取)
# ・機械学習
# ・上の手法の組み合わせ

# シュミレーションの流れ
# Validators -> Committee -> Vote

# 正しいブロックを作れるかについての実験
# 状況
# validators n人 pは正規分布
# ブロックチェーンがある

import random
import numpy as np

# パラメータ
alpha = 0.5
delta = 0.001
la = 12
lr = 0.01
q = 2 / 3
m = 15

c_alpha = 1
c_beta = 3
c_delta = 1
c_gamma = 1

class Block:
    def __init__(self, created_by, staking_token, age, reputation, random):
        self.created_by = created_by
        self.staking_token = staking_token
        self.age = age
        self.reputation = reputation
        self.random = random

def create_block_chain(n, validators):
    blockchain = []
    vs_count = len(validators)
    for i in range(n):
        proposed_block = 0 if alpha > random.random() else 1
        winner = winner_by_graph(create_graph(validators),validators)
        winner_decision = winner.p_to_decision(proposed_block)
        if winner_decision == proposed_block and proposed_block == 1:
            block = Block(winner.name, winner.staking_token, winner.age, winner.reputation, winner.random)
            blockchain.append(block)
            winner.staking_token += 1000
            winner.reputation += 1
            for validator in validators:
                validator.age += 1
            winner.age = 1
        elif winner_decision != proposed_block:
            add_stake = winner.staking_token / vs_count
            for validator in validators:
                validator.staking_token += add_stake
            winner.staking_token = 0
        winner.update_profile(winner_decision, proposed_block, delta, lr, la)
    return blockchain

# いい感じにvalidatorsとblockchainを初期化してくれる
def setup_simulation():
    vs = create_validators(0.75, 0.1, 100, 0)
    bc = create_block_chain(1200, vs)
    return {"validators": vs, "blockchain": bc}

def attack_validators(vs, rate):
    attacked_num = int(len(vs) * rate)
    for v in vs:
        v.attacked = False
    for v in random.sample(vs, attacked_num):
        v.attacked = True

# パラメータ
alpha = 0.5
delta = 0.001
la = 12
lr = 0.01
q = 2 / 3
m = 15

c_alpha = 1
c_beta = 3
c_delta = 1
c_gamma = 1

# vs_and_blockchain = setup_simulation()
# vs = vs_and_blockchain["validators"]
# blockchain = vs_and_blockchain["blockchain"]
# attack_validators(vs, 0.25)

c_beta = 2 # セットアップのために3にしていたが戻す

wmv_ecws = []
mv_ecws = []
mvw_ecw = []
rd_ecw = []
cd_ecw = []

vs_count = len(vs)
for t in range(100):
    print(t)
    proposed_block = 0 if alpha > random.random() else 1

    xs = [v.p_to_decision(proposed_block) for v in vs]
    ws = [v.weight() for v in vs]

    mv_ecws.append(ecw(pav(vs, majority_rules, 50), pri(vs, majority_rules, 50)))
    mvw_ecw.append(ecw(pav(vs, m_votes_to_win, 50), pri(vs, m_votes_to_win, 50)))
    wmv_ecws.append(ecw(pav(vs, weighted_majority_rule, 50), pri(vs, weighted_majority_rule, 50)))
    rd_ecw.append(ecw(pav(vs, random_decision, 50), pri(vs, random_decision, 50)))
    cd_ecw.append(ecw(pav(vs, condrcet_dicision,50), pri(vs, condrcet_dicision, 50)))

    # profileを更新する
    for i in range(len(vs)):
        vs[i].update_profile(xs[i], proposed_block, delta, lr, la)

    # コンドルセのパラメータも更新する
    winner = winner_by_graph(create_graph(vs),vs)
    winner_decision = winner.p_to_decision(proposed_block)
    if winner_decision == proposed_block and proposed_block == 1:
        block = Block(winner.name, winner.staking_token, winner.age, winner.reputation, winner.random)
        blockchain.append(block)
        winner.staking_token += 1000
        winner.reputation += 1
        for validator in vs:
            validator.age += 1
        winner.age = 1
    elif winner_decision != proposed_block:
        add_stake = winner.staking_token / vs_count
        for validator in vs:
            validator.staking_token += add_stake
        winner.staking_token = 0

x_axis = list(range(100))
plt.title('weighted approval votes')
plt.xlabel('times')
plt.ylabel('ecw')
plt.plot(x_axis, mv_ecws, label="mv")
plt.plot(x_axis, mvw_ecw, label="mvw")
plt.plot(x_axis, wmv_ecws, label="wmv")
plt.plot(x_axis, rd_ecw, label="rd")
plt.plot(x_axis, cd_ecw, label="cd")
plt.legend()
plt.savefig("plot_ecw_by_times.png")
plt.show()
