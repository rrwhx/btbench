import time
from aes import AESModeOfOperation


conversionFunc = lambda x: x*1000000


if __name__ == "__main__":

    print("\n")
    print("				#################")
    print("				####benchmark####")
    print("				#################\n")
    
    
    beforetotal=time.clock()

    for i in range(30):
        # before=time.clock()
        print("\n")
        print("Measuring test AESBench iteration", i + 1)
        AESModeOfOperation().start()
        # after=time.clock()
        # execution_time=int(conversionFunc(after-before))
        # print("AESBench;",execution_time)
    
    aftertotal=time.clock()
    execution_time=float((aftertotal-beforetotal))
    print("total execution time;",execution_time, "s")    
    
    
    
    
    
    
    
