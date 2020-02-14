import random
# パラメータ
alpha = 0.5
la = 12
lr = 0.01

# 社会余剰
# expected colective welfare -> maximize
def ecw(p):
    return (1 - alpha) * (1 + lr) * p + alpha * (1 + la) * p

def agg_p(x,y):
    return x + y - x*y
