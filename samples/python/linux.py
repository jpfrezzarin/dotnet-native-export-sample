from ctypes import cdll

print("Hello, World!")

# give location of dll
mydll = cdll.LoadLibrary("output/MyNativeLibraryNE.so")

sumResult = mydll.MyLibSum(1,2)
subResult = mydll.MyLibSub(4,2)

print("The result of sum is: " + str(sumResult))
print("The result of sub is: " + str(subResult))