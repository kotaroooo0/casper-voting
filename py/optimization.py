# Committee size = 3
#
validators = [0.9, 0.9, 0.9, 0.6, 0.6, 0.6]
max_p = -1
ret = []
for i in range(len(validators)):
    for j in range(len(validators)):
        for k in range(len(validators)):
            if (i == j or j == k or k == i) {
                continue
            }
            vp = validators[i]*validators[j] + validators[i]*validators[j] + \
                validators[i]*validators[j] - 2 * \
                validators[i]*validators[j]*validators[k]
            if vp >= max_p:
                max_p = vp
                ret[0] = max_p
                ret.append([i, j, k])
