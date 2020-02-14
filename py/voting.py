import seaborn as sns
import matplotlib.pyplot as plt
import random
import numpy as np
from statistics import mean, variance

%matplotlib inline

# 効用に関するもの

# 社会余剰
# expected colective welfare -> maximize


def ecw(pav, pri):
    return (1 - alpha) * (1 + lr) * pav + alpha * (1 + la) * pri

# the probabilities of accepting a valid block


def pav(validators, oppinion_aggregation_method, ite=10000):
    opinions = [oppinion_aggregation_method(validators, 1) for _ in range(ite)]
    return opinions.count(1) / ite

# the probabilities　of rejecting an invalid block


def pri(validators, oppinion_aggregation_method, ite=10000):
    opinions = [oppinion_aggregation_method(
        validators, -1) for _ in range(ite)]
    return opinions.count(-1) / ite

# 意見集約法に関わるもの
# decision rule

# 多数決


def majority_rules(validators, proposed_block):
    xs = list(map(lambda v: v.p_to_decision(proposed_block), validators))
    if(sum(xs) >= (2*q - 1)*len(validators)):
        return 1
    else:
        return -1

# m票先取


def m_votes_to_win(validators, proposed_block):
    agree_votes_count = 0
    disagree_votes_count = 0
    for validator in validators:
        if validator.p_to_decision(proposed_block) == 1:
            agree_votes_count += 1
            if agree_votes_count >= m:
                return 1
        else:
            disagree_votes_count += 1
            if disagree_votes_count >= m:
                return -1
    return 0
    raise Exception("Processing did not end")

# ランダム(最初の一票)


def random_decision(validators, proposed_block):
    return random.choice(validators).p_to_decision(proposed_block)

# シンプルなコンドルセ


def condrcet_dicision(validators, proposed_block):
    winner = winner_by_graph(create_graph(validators), validators)
    return winner.p_to_decision(proposed_block)

# 重み付き多数決


def weighted_majority_rule(validators, proposed_block):
    ws = list(map(lambda v: v.weight(), validators))
    xs = list(map(lambda v: v.p_to_decision(proposed_block), validators))
    if(np.dot(ws, xs) >= (2*q - 1)*sum(ws)):
        return 1
    else:
        return -1


def plot_ecw_mr(vs):
    x_axis = np.linspace(0.51, 1, 50)
    y_axis = []
    for x in x_axis:
        global q
        q = x
        y_axis.append(ecw(pav(vs, majority_rules), pri(vs, majority_rules)))
    plt.title('ecw mr alpha: {0} la: {1} lr: {2}'.format(alpha, la, lr))
    plt.xlabel('q')
    plt.ylabel('ecw')
    plt.plot(x_axis, y_axis)
    plt.savefig("plot_ecw_mr.png")
    plt.show()


def plot_ecw_mvw(vs):
    x_axis = np.linspace(1, len(vs), len(vs))
    y_axis = []
    for x in x_axis:
        global m
        m = x
        y_axis.append(ecw(pav(vs, m_votes_to_win), pri(vs, m_votes_to_win)))
    plt.title('ecw mvw alpha: {0} la: {1} lr: {2}'.format(alpha, la, lr))
    plt.xlabel('m')
    plt.ylabel('ecw')
    plt.plot(x_axis, y_axis)
    plt.savefig("plot_ecw_mvw.png")
    plt.show()


def ecw_mean_and_val(vs, oppinion_aggregation_method, ite):
    ecws = [ecw(pav(vs, oppinion_aggregation_method), pri(
        vs, oppinion_aggregation_method)) for x in range(ite)]
    return [mean(ecws), variance(ecws)]
