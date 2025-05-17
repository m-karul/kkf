
#define openfd[int32] =0 
#define t[int64] =2 
#define name[char] ="Ali"


struct Person {
    char name[128] <k ==2 ? 2:1>
    int32 age 
}
 
struct Address {
    int32 city  
    int32 zip 
}

#Flags flag1=AAA,BBB,CCC,DD
#Flags flag2=AAA,BBB,CCC,DD

#   özellikler:
#       range:1000-100000
#       flags:flag_degisken_adi
#       const:flag adı yada değeri
#       
#

$open@first_open(Const:"./file", Range:int32, Flags:flag1, Struct:Person) -> openfd
$ioctl@aaa(Const:"./file2", Range:1-200, Flags:flag1, Const:"AA") -> openfd

