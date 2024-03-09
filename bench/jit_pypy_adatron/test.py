import time
import adatron


conversionFunc = lambda x: x*1000000


if __name__ == "__main__":

    print("\n")
    print("				##################")
    print("				#####programs#####")
    print("				##################\n")
    
    beforetotal=time.clock()
    
    for i in range(3):
        before=time.clock()
        print("\n")
        print("Measuring test AdatronSVM iteration", i + 1)
        adatron.run()
        after=time.clock()
        execution_time=int(conversionFunc(after-before))
        # print("AdatronSVM;",execution_time)
    
    aftertotal=time.clock()
    execution_time=float((aftertotal-beforetotal))
    print("total execution time:",execution_time, "s")    
    
    
    
    
    
    
    
