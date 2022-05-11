from ctypes import*

# give location of dll
mydll = cdll.LoadLibrary(".\\..\\..\\output\\MyNativeLibraryNE.dll")

sumResult = mydll.Sum(1,2)
subResult = mydll.Sub(4,2)

print("The result of sum is: " + str(sumResult))
print("The result of sub is: " + str(subResult))