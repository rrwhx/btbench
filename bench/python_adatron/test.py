import time
import adatron


conversionFunc = lambda x: x*1000000


if __name__ == "__main__":

    print("\n")
    print("				##################")
    print("				#####programs#####")
    print("				##################\n")
    
    beforetotal=time.perf_counter_ns()
    
    for i in range(1):
        before=time.perf_counter_ns()
        print("\n")
        print("Measuring test AdatronSVM iteration", i + 1)
        adatron.run()
        after=time.perf_counter_ns()
        execution_time=int(conversionFunc(after-before))
        # print("AdatronSVM;",execution_time)
    
    aftertotal=time.perf_counter_ns()
    execution_time=float((aftertotal-beforetotal)/1000000000.0)
    print("total execution time:",execution_time, "s")    
    
    
    
    
    
    
    
