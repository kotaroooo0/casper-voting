# import numpy as np
import copy
from statistics import variance

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

# validatorsとblockchainを準備
# validators100人
# vs_and_blockchain = setup_simulation()
vs = vs_and_blockchain["validators"]
blockchain = vs_and_blockchain["blockchain"]
vs_for_staking_token = copy.deepcopy(vs)
vs_for_all = copy.deepcopy(vs)
vs_for_condrcet = copy.deepcopy(vs)

committee_size = 11

times = 500
ite = 10
sc_variance = [0]*times
ac_variance = [0]*times
cc_variance = [0]*times
sc_ecw = [0]*times
ac_ecw = [0]*times
for it in range(ite):
    vs_for_staking_token = copy.deepcopy(vs)
    vs_for_all = copy.deepcopy(vs)
    vs_for_condrcet = copy.deepcopy(vs)
    temp_vs = copy.copy(vs_for_condrcet)
    for t in range(times):
        proposed_block = 0 if alpha > random.random() else 1

        # staking_tokenの保有率に応じたcommitteeを作成
        sum_staking_token = sum([v.staking_token for v in vs_for_staking_token])
        staking_committee = np.random.choice(vs_for_staking_token, size=committee_size, p=[v.staking_token / sum_staking_token for v in vs_for_staking_token], replace=False)

        scxs = [v.p_to_decision(proposed_block) for v in staking_committee]
        committee_decision = -1
        if(sum(scxs) >= (2*q - 1)*committee_size): committee_decision = 1

        # 新規通貨は1000
        # 間違えた人は没収され、正しい答えの人に分配される
        destributed_token = 1000
        correct_validators_count = 0
        for i in range(committee_size):
            if (scxs[i] == proposed_block):
                correct_validators_count += 1
            else:
                destributed_token += staking_committee[i].staking_token * 0.8
                staking_committee[i].staking_token = staking_committee[i].staking_token * 0.2
        for i in range(committee_size):
            if (scxs[i] == proposed_block):
                staking_committee[i].staking_token += destributed_token / correct_validators_count

        sc_variance[t] += variance([v.staking_token for v in vs_for_staking_token]) / ite
        sc_ecw[t] += ecw(pav(staking_committee, majority_rules, 100), pri(staking_committee, majority_rules, 100)) / ite

        # age, reputation, randomも考慮したcommitteeを作成
        sum_age = sum([v.age for v in vs_for_all])
        sum_reputation = sum([v.reputation for v in vs_for_all])
        sum_random = sum([v.random for v in vs_for_all])
        all_committee = np.random.choice(vs_for_all, size=committee_size, p=[(v.staking_token / sum_staking_token + v.age / sum_age + v.reputation / sum_reputation + v.random / sum_random) / 4 for v in vs_for_all], replace=False)

        acvs = [v.p_to_decision(proposed_block) for v in all_committee]

        committee_decision = -1
        if(sum(acvs) >= (2*q - 1)*committee_size): committee_decision = 1

        # 新規通貨は1000
        # 間違えた人は没収され、正しい答えの人に分配される
        destributed_token = 1000
        correct_validators_count = 0
        for i in range(committee_size):
            if (acvs[i] == proposed_block):
                correct_validators_count += 1
            else:
                destributed_token += all_committee[i].staking_token * 0.8
                all_committee[i].staking_token = all_committee[i].staking_token * 0.2
        for i in range(committee_size):
            if (acvs[i] == proposed_block):
                all_committee[i].staking_token += destributed_token / correct_validators_count

        ac_variance[t] += variance([v.staking_token for v in vs_for_all]) / ite
        ac_ecw[t] += ecw(pav(all_committee, majority_rules, 100), pri(all_committee, majority_rules, 100)) / ite

        # condrcetなcommitteeを作成
#         condrcet_committee = []
#         for _ in range(committee_size):
#             winner = winner_by_graph(create_graph(temp_vs),temp_vs)
#             condrcet_committee.append(winner)
#             temp_vs.remove(winner)

#         ccxs = [v.p_to_decision(proposed_block) for v in condrcet_committee]
#         committee_decision = -1
#         if(sum(ccxs) >= (2*q - 1)*committee_size): committee_decision = 1

        # 新規通貨は1000
        # 間違えた人は没収され、正しい答えの人に分配される
#         destributed_token = 1000
#         correct_validators_count = 0
#         for i in range(committee_size):
#             if (ccxs[i] == proposed_block):
#                 correct_validators_count += 1
#             else:
#                 destributed_token += condrcet_committee[i].staking_token * 0.8
#                 condrcet_committee[i].staking_token = condrcet_committee[i].staking_token * 0.2
#         for i in range(committee_size):
#             if (scxs[i] == proposed_block):
#                 condrcet_committee[i].staking_token += destributed_token / correct_validators_count

#         cc_variance[t] += variance([v.staking_token / ite for v in vs_for_condrcet])
#         print(variance([v.staking_token for v in vs_for_condrcet]))
#         print(t)

# print(variance([v.staking_token for v in staking_committee]))
# print(variance([v.staking_token for v in all_committee]))
# print(variance([v.staking_token for v in vs]))

x_axis = list(range(times))
plt.title('staking token variance by committee')
plt.xlabel('times')
plt.ylabel('ecw')
plt.plot(x_axis, sc_variance, label="committee by staking token")
plt.plot(x_axis, ac_variance, label="committee by all parameters")
plt.legend()
plt.savefig("plot_staking_token_variance_by_committee.png")
plt.show()

plt.title('ecw by committee')
plt.xlabel('times')
plt.ylabel('ecw')
plt.plot(x_axis, sc_ecw, label="committee by staking token")
plt.plot(x_axis, ac_ecw, label="committee by all parameters")
plt.legend()
plt.savefig("plot_ecw_by_committee.png")
plt.show()
