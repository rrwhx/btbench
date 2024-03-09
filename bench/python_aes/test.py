import time
from aes import AESModeOfOperation


conversionFunc = lambda x: x*1000000


if __name__ == "__main__":

    print("\n")
    print("				#################")
    print("				####benchmark####")
    print("				#################\n")
    
    
    beforetotal=time.perf_counter_ns()

    for i in range(2):
        # before=time.perf_counter_ns()
        print("\n")
        print("Measuring test AESBench iteration", i + 1)
        AESModeOfOperation().start()
        # after=time.perf_counter_ns()
        # execution_time=int(conversionFunc(after-before))
        # print("AESBench;",execution_time)
    
    aftertotal=time.perf_counter_ns()
    execution_time=float((aftertotal-beforetotal)/1000000000.0)
    print("total execution time;",execution_time, "s")    
    
    
    
    
    
    
    
