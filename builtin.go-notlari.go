package builtin

import "cmp"

type bool bool

const (
	true  = 0 == 0 // Tip tanımlamaları için boolean değerler true ve false.
	false = 0 != 0 // Tip tanımlamaları için boolean değerler true ve false.
)

type uint8 uint8 // Tüm 8-bitlik işaretsiz tamsayıları temsil eder. Aralık: 0 ile 255.

type uint16 uint16 // Tüm 16-bitlik işaretsiz tamsayıları temsil eder. Aralık: 0 ile 65535.

type uint32 uint32 // Tüm 32-bitlik işaretsiz tamsayıları temsil eder. Aralık: 0 ile 4294967295.

type uint64 uint64 // Tüm 64-bitlik işaretsiz tamsayıları temsil eder. Aralık: 0 ile 18446744073709551615.

type int8 int8 // Tüm 8-bitlik işaretli tamsayıları temsil eder. Aralık: -128 ile 127.

type int16 int16 // Tüm 16-bitlik işaretli tamsayıları temsil eder. Aralık: -32768 ile 32767.

type int32 int32 // Tüm 32-bitlik işaretli tamsayıları temsil eder. Aralık: -2147483648 ile 2147483647.

type int64 int64 // Tüm 64-bitlik işaretli tamsayıları temsil eder. Aralık: -9223372036854775808 ile 9223372036854775807.

type float32 float32 // Tüm IEEE-754 32-bit kayan noktalı sayıları temsil eder.

type float64 float64 // Tüm IEEE-754 64-bit kayan noktalı sayıları temsil eder.

type complex64 complex64 // float32 gerçel ve sanal kısımlara sahip tüm karmaşık sayıları temsil eder.

type complex128 complex128 // float64 gerçel ve sanal kısımlara sahip tüm karmaşık sayıları temsil eder.

type string string // UTF-8 kodlu metni temsil eder. Boş bir metin olabilir, ancak nil olamaz. string türündeki değerler değiştirilemez.

type int int // En az 32 bit boyutunda işaretli bir tamsayı türüdür. Aynı zamanda int32 gibi bir takma ad değildir.

type uint uint // En az 32 bit boyutunda işaretsiz bir tamsayı türüdür. Aynı zamanda uint32 gibi bir takma ad değildir.

type uintptr uintptr // İşaretçi verilerin bit desenini tutmak için yeterli büyüklükte bir tamsayı türüdür.

type byte = uint8 // uint8 ile eşdeğer olan ve 8-bitlik işaretsiz tamsayı değerlerini temsil eden bir takma adır.

type rune = int32 // int32 ile eşdeğer olan ve karakter değerlerini temsil eden bir takma adır.

type any = interface{} // interface{} ile eşdeğer olan ve herhangi bir Go tipini temsil eden bir takma adır.

type comparable interface{ comparable } // Tüm karşılaştırılabilir tipler tarafından uygulanan bir arayüzdür (booleanlar, sayılar, dizeler, işaretçiler, karşılaştırılabilir türlerden oluşan diziler, tüm alanları karşılaştırılabilir tiplerden oluşan yapılar).

const iota = 0 // Tip tanımlamaları içinde geçerli olan geçerli const belirtecinin sıfırdan başlayan sırasal sayısıdır.

var nil Type // Type türünden bir pointer, channel, func, interface, map veya slice türü için sıfır değerini temsil eden önceden tanımlanmış bir tanımlayıcı.

type Type int // Type, belirli bir işlev çağrısındaki aynı türü temsil eden, yalnızca belgeleme amacıyla burada bulunan bir yer tutucudur.

type Type1 int // Type1, belirli bir işlev çağrısındaki aynı türü temsil eden, yalnızca belgeleme amacıyla burada bulunan bir yer tutucudur.

type IntegerType int // IntegerType, herhangi bir tamsayı türünü temsil eden, yalnızca belgeleme amacıyla burada bulunan bir yer tutucudur (örneğin: int, uint, int8 vb.).

type FloatType float32 // FloatType, float32 veya float64 türünden herhangi bir ondalık sayı türünü temsil eden, yalnızca belgeleme amacıyla burada bulunan bir yer tutucudur.

type ComplexType complex64 // ComplexType, complex64 veya complex128 türünden herhangi bir karmaşık sayı türünü temsil eden, yalnızca belgeleme amacıyla burada bulunan bir yer tutucudur.

func append(slice []Type, elems ...Type) []Type // Bir dilim sonuna öğeler ekler. Yeterli kapasitesi varsa hedef, yeni öğeleri barındıracak şekilde tekrar kesilir. Eğer kapasitesi yoksa, yeni bir altta yatan dizi ayrılır. append işlemi güncellenmiş dilimi döndürür. Bu nedenle genellikle dilimi tutan değişkende sonucu saklamak gerekir.

func copy(dst, src []Type) int // copy, kaynaktan hedefe öğeleri kopyalar. (Aynı zamanda bir dizeden byte dilimine de kopyalama yapar.) Kaynak ve hedef üst üste gelebilir. copy işlemi, kopyalanan öğe sayısını döndürür; bu sayı, len(src) ile len(dst) değerlerinin minimumudur.

func delete(m map[Type]Type1, key Type) // delete, haritadan belirtilen anahtarı (m[anahtar]) siler. Eğer m nil ise veya böyle bir öğe yoksa, delete işlemi bir işlem yapmaz.

func len(v Type) int // len, türüne göre v'nin uzunluğunu döndürür:
//	Dizi: v'deki öğe sayısı.
//	Dizi işaretçisi: *v'deki öğe sayısı (v nil ise bile).
//	Dilim veya harita: v'deki öğe sayısı; eğer v nil ise, len(v) sıfırdır.
//	Dize: v'deki byte sayısı.
//	Kanal: kanal tamponunda (okunmamış) öğe sayısı;
//	v nil ise, len(v) sıfırdır.
//	Bazı argümanlar için, örneğin bir dize harfesi veya basit bir dizi ifadesi için, sonuç bir sabit olabilir.

func cap(v Type) int // cap, türüne göre v'nin kapasitesini döndürür:
//	Dizi: v'deki öğe sayısı (len(v) ile aynı).
//	Dizi işaretçisi: *v'deki öğe sayısı (len(v) ile aynı).
//	Dilim: dilim tekrar kesildiğinde ulaşabileceği maksimum uzunluk;
//	v nil ise, cap(v) sıfırdır.
//	Kanal: öğe birimleri cinsinden kanal tampon kapasitesi;
//	v nil ise, cap(v) sıfırdır.
//	Bazı argümanlar için, örneğin basit bir dizi ifadesi için, sonuç bir sabit olabilir.

func make(t Type, size ...IntegerType) Type // make, bir dilim, harita veya kanal (yalnızca) için bellek ayırır ve başlatır. new gibi, ilk argüman bir türdür, bir değil. Ancak make'nin dönüş türü, türün bir işaretçisi değil, kendisidir. Sonuç türünün belirtimi türün türüne bağlıdır:
//	Dilim: Boyut, dilim uzunluğunu belirtir. Dilimin kapasitesi uzunluğuna eşittir. Farklı bir kapasite belirtmek için ikinci bir tamsayı argümanı sağlanabilir; bu, uzunluktan küçük olmamalıdır. Örneğin make([]int, 0, 10), boyutu 10 olan bir temel dizin ayrır ve uzunluğu 0 ve kapasitesi 10 olan bir dilim döndürür.
//	Harita: Belirtilen öğe sayısını barındırmak için yeterli alanla boş bir harita ayrılır. Boyut belirtilebilir, bu durumda küçük bir başlangıç boyutu ayrılır.
//	Kanal: Kanalın tamponu belirtilen tampon kapasitesiyle başlatılır. Sıfır veya boyut belirtilmezse, kanal tamponlu değildir.

func max[T cmp.Ordered](x T, y ...T) T // max, bir sabit sayıda [cmp.Ordered] tipten argümanın en büyük değerini döndürür. En az bir argüman olmalıdır. T, bir kayan nokta türü ise ve argümanlardan herhangi biri NaN ise, max NaN döndürür.

func min[T cmp.Ordered](x T, y ...T) T // min, bir sabit sayıda [cmp.Ordered] tipten argümanın en küçük değerini döndürür. En az bir argüman olmalıdır. T, bir kayan nokta türü ise ve argümanlardan herhangi biri NaN ise, min NaN döndürür.

func new(Type) *Type // new, bellek ayırır. İlk argüman bir türdür, bir değil; ve dönen değer yeni bir türün işaretçisidir ve değil.

func complex(r, i FloatType) ComplexType // complex, iki kayan nokta değerinden karmaşık bir değer oluşturur. Gerçek ve sanal kısımlar aynı boyutta olmalıdır, ya float32 ya da float64 (veya bunlara atanabilir). Dönüş değeri, karşılık gelen karmaşık türdür (float32 için complex64, float64 için complex128).

func real(c ComplexType) FloatType // real, karmaşık sayının gerçek kısmını döndürür. Dönüş değeri, c'nin türüne uygun olan ondalık sayı türüdür.

func imag(c ComplexType) FloatType // imag, karmaşık sayının sanal kısmını döndürür. Dönüş değeri, c'nin türüne uygun olan ondalık sayı türüdür.

func clear[T ~[]Type | ~map[Type]Type1](t T) // clear, haritaları ve dilimleri temizler. Haritalar için, tüm girdileri siler, sonuç olarak boş bir harita elde edilir. Dilimler için, tüm öğeleri sıfır değerine ayarlar. Eğer argüman tür parametresi ise, tür parametresinin tür kümesi yalnızca harita veya dilim türlerini içermeli ve clear işlemi tür argümanının belirttiği işlemi gerçekleştirir.

func close(c chan<- Type) // close, bidirectional veya sadece gönderme yapılabilen bir kanalı kapatır. Bu, gönderici tarafından yürütülmelidir, alıcı tarafından değil. close işlemi, son gönderilen değer alındıktan sonra kanalı kapatmanın etkisini yapar. Kapatılan bir kanaldan alındığında her zaman bir değer döner; bu değer, kanal öğesi için sıfır değerdir. Kapatılmış ve boş bir kanaldan alındığında:
//	x, ok := <-c
// ok, kapalı ve boş bir kanal için false olarak ayarlanır.

func panic(v any) // panic, mevcut gorutinin normal yürütmesini durdurur. Bir işlev F'nin panic çağrısı yaptığı yerde F'nin normal yürütmesi hemen durur. F tarafından ertelenen işlevlerin yürütümü normal bir şekilde gerçekleştirilir ve ardından F çağrıcısına döner. Çağrıcı G için, F'nin çağrısı bir panic çağrısı gibi davranır, G'nin yürütmesini sonlandırır ve ertelenen işlevleri çalıştırır. Bu, yürütülen gorutinin tüm işlevleri ters sırayla durduğu noktada gerçekleşir. Bu noktada, program, çıkış kodu sıfır olmayan bir şekilde sonlandırılır. Bu sonlandırma dizisi paniklemek olarak adlandırılır ve built-in fonksiyon recover tarafından kontrol edilebilir.
// Go 1.21'den itibaren, panic' in nil arayüz değeri veya tipi belirtilmemiş nil ile çağrılması bir çalışma zamanı hatası oluşturur (farklı bir panik).
// GODEBUG ayarı panicnil=1 ise bu çalışma zamanı hatasını devre dışı bırakır.

func recover() any // recover, bir programın panik yapma durumunu yönetmesine olanak tanır. recover'ın bir ertelenen işlev içinde çağrılması (ancak onun tarafından çağrılan işlevlerde değil) panik sırasını durdurur, normal yürütmeyi geri yükler ve panik çağrısına iletilen hata değerini alır. recover, ertelenen işlev dışında çağrıldığında panik sırasını durdurmaz. Bu durumda veya gorutin paniklemiyorsa veya panic'e nil bir argüman sağlanmışsa, recover nil döner. Bu nedenle, recover'ın dönüş değeri gorutinin paniklemesi durumunu bildirir.

func print(args ...Type) // print, argümanlarını uygulama özgü bir şekilde biçimlendirir ve sonucu standart hataya yazar. Print, önyükleme yapmak ve hata ayıklamak için kullanışlıdır; dilimde kalmayacağı garanti edilmez.

func println(args ...Type) // println, argümanlarını uygulama özgü bir şekilde biçimlendirir ve sonucu standart hataya yazar. Arga sıra eklenir ve bir satır sonlandırılır. Println, önyükleme yapmak ve hata ayıklamak için kullanışlıdır; dilimde kalmayacağı garanti edilmez.

type error interface {
	Error() string
}
