# 重み付き投票のシュミレーション
import random

%matplotlib inline
import matplotlib.pyplot as plt

# パラメータ
alpha = 0.5
delta = 0.001
la = 12
lr = 0.01
q = 2 / 3
m = 10

# return: 攻撃されていないvalidatorの重み / 全てのvalidatorの重み
def weighted_approval_votes(validators):
    non_attacked_validator_weight = 0
    all_validator_weight = 0
    for validator in validators:
        w = validator.weight()
        all_validator_weight += w
        if not validator.attacked:
            non_attacked_validator_weight += w
    return non_attacked_validator_weight / all_validator_weight

def weighted_approval_votes_changes(times, validators):
    result = []
    for t in range(times):
        result.append(weighted_approval_votes(validators))

        # 提案されたブロックを定義
        proposed_block = 0 if alpha > random.random() else 1

        # 実際に投票を行いprofileを更新する
        for validator in validators:

            # 攻撃されている場合、投票することができない
            if validator.attacked: continue

            if validator.p > random.random():
                # 正しく投票のとき
                validator.update_profile(proposed_block, proposed_block, delta, lr, la)
            else:
                # 間違った投票のとき
                validator.update_profile(-proposed_block, proposed_block, delta, lr, la)
    return result


def plot_weighted_approval_votes(times, validators):
    x_axis = list(range(times))
    y_axis = weighted_approval_votes_changes(times, validators)
    plt.title('weighted approval votes')
    plt.xlabel('times')
    plt.ylabel('weighted approval votes')
    plt.plot(x_axis, y_axis)
    plt.savefig("plot_weighted_approval_votes.png")
    plt.show()
