from matplotlib import pyplot as plt
from math import log2

x1 = [8*(i**2) for i in range(1,100)]
x2 = [64*i*log2(i) for i in range(1,100)]

plt.plot(x1, label="8n^2")
plt.plot(x2, label="64nlg(n)")
plt.axvline(x=43, color='gray', linestyle='--', label="intersect")
plt.legend()
plt.xlabel("Size of n")
plt.ylabel("Runtime")
plt.title("Order of growth intersection")
plt.show()

