from Crypto import crypt, convert


def encrypt(x):
    m, pad = convert.t_to_a(x)
    print(x, m, "\n\n")
    c, d, n = crypt.encrypt(m)
    print("\n", c, "\n")

def decrypt(m):
