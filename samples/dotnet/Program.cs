using DotnetSample;

Console.WriteLine("Hello, World!");

int sum = MyNativeLibrary.MyLibSum(1, 2);
int sub = MyNativeLibrary.MyLibSub(3, 2);

Console.WriteLine($"The result of sum is: {sum}");
Console.WriteLine($"The result of sub is: {sub}");
