import math
import random
import string

class Validator():
    def __init__(self, p, name, attacked, staking_token, age, reputation, random):
        # profile
        self.p = p

        # バリデータを識別する名前
        # 数値シュミレーションにおいて意味はない
        self.name = name

        # 攻撃されているかどうか
        # 攻撃されているバリデータは投票できない
        self.attacked = attacked

        # Refer: BLOCKCHAIN: A NOVEL APPROACH FOR THE CONSENSUS ALGORITHM USING CONDORCET VOTING PROCEDURE
        self.staking_token = staking_token
        self.age = age
        self.reputation = reputation
        self.random = random

    def update_profile(self, x, b, delta, lr, la):
        if x == b:
            self.p = min(self.p * (1+delta), 0.99999)
        elif b == 1:
            self.p = max(self.p * ((1-delta)**lr), 0.5)
        else:
            self.p = max(self.p * ((1-delta)**la), 0.5)

    def weight(self):
        return math.log(self.p) - math.log(1.0-self.p)

    def p_to_decision(self, proposed_block):
        if self.attacked: return -1
        return proposed_block if self.p > random.random() else -1 * proposed_block

# バリデータ
# 正答率が正規乱数に従うn人のバリデータを生成
def create_validators(mean, var, n, attacked_rate):
    attacked_validator_count = int(n * attacked_rate)
    aps = np.random.normal(mean, var, attacked_validator_count).tolist()
    attacked_validators = list(map(lambda p: Validator(max(min(0.99999, p), 0.50001), randomname(), True, 1000, 0, 0, random.random()), aps))
    non_attacked_validator_count = n - attacked_validator_count
    nps = np.random.normal(mean, var, non_attacked_validator_count).tolist()
    non_attacked_validators = list(map(lambda p: Validator(max(min(0.99999, p), 0.50001), randomname(), False, 1000, 0, 0, random.random()), nps))
    validators = attacked_validators + non_attacked_validators
    random.shuffle(validators)
    return validators

def randomname():
    return ''.join(random.choices(string.ascii_letters + string.digits, k=10))

def setup_init_validators(validators, blockchain):
    set_init_staking_token(validators)
    set_init_age(validators, blockchain)
    set_init_reputation(validators, blockchain)

# validatorのstaking token の初期値を1000とする
def set_init_staking_token(validators):
    for validator in validators:
        validator.staking_token = 1000

# 最後に作ったブロックの深さ
def set_init_age(validators, blockchain):
    blockchain.reverse()
    for validator in validators:
        isDetermined = False
        for i in range(len(blockchain)):
            if (blockchain[i].created_by == validator.name):
                validator.age = i+1
                isDetermined = True
                break
        if not isDetermined: validator.age = len(blockchain)

# 今までで作ったブロック数
def set_init_reputation(validators, blockchain):
    for validator in validators:
        rep = 0
        for block in blockchain:
            if (block.created_by == validator.name):
                rep += 1
        validator.reputation = rep
