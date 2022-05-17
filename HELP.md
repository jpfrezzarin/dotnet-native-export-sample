# Help

## Possible errors

### 1. GLIBCXX not found

If you try to import the shared object (.so) in some application in Linux and get the error bellow, it's probably because you don't have GLIBCXX.

```
libstdc++.so.6: version `GLIBCXX_3.4.20' not found
```

To install it, run the commands bellow in terminal:

```
sudo add-apt-repository ppa:ubuntu-toolchain-r/test
sudo apt-get update
sudo apt-get install gcc-4.9
sudo apt-get upgrade libstdc++6
```