import voting
import validator
import pytest

import numpy as np

def test_validators():
    validators = voting.validators(70, 5, 10)

    for i in range(1,10):
        print (i)
        print (validators[i].p)

def test_majority_rules():
    x = np.array([1, 1, 1, 1])
    q = 0.5
    n = x.size
    assert voting.majority_rules(x, q, n) == 1

    x = np.array([1, 1, -1, -1])
    q = 0.5
    n = x.size
    assert voting.majority_rules(x, q, n) == 1

    x = np.array([-1, -1, -1, 1])
    q = 0.5
    n = x.size
    assert voting.majority_rules(x, q, n) == -1

def test_m_votes_to_win():
    x = np.array([-1, 1, -1, -1])
    m = 3
    assert voting.m_votes_to_win(x, m) == -1

    x = np.array([1, 1, -1, 1])
    m = 3
    assert voting.m_votes_to_win(x, m) == 1

    x = np.array([1, 1, -1, 1])
    m = 4
    assert voting.m_votes_to_win(x, m) == 0

    with pytest.raises(Exception):
            x = np.array([1, 11, -1, 1])
            m = 4
            voting.m_votes_to_win(x, m)

def test_random_decision():
    x = np.array([1, 1, -1, 1])
    assert voting.random_decision(x) == 1

    x = np.array([-1, 1, -1, 1])
    assert voting.random_decision(x) == -1

    with pytest.raises(Exception):
        x = np.array([121, 1, -1, 1])
        assert voting.random_decision(x) == 1
