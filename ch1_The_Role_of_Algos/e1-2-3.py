from matplotlib import pyplot as plt

def poly(n):
    return 100*n**2

def expo(n):
    return 2**n

n = 1
while poly(n) > expo(n):
    n += 1
n += 1



x1 = [poly(i) for i in range(n)]
x2 = [expo(i) for i in range(n)]

plt.plot(x1, label="100n^2")
plt.plot(x2, label="2^n")
plt.legend()
plt.xlabel("Size of n")
plt.ylabel("Runtime")
plt.title("Order of growth intersection")
plt.show()