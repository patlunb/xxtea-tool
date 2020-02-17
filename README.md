# xxtea-tool
**XXTEA Encrypt and Decrypt Command Line Tool**

Simple command-line interface for the xxtea algorithm string encryption and decryption.

 **Based on [hillu](https://github.com/hillu/go-xxtea "hillu's") ** library, that uses newest implementation [updated](http://www.movable-type.co.uk/scripts/xxtea.pdf "updated") by *Needham and Wheeler* as in https://en.wikipedia.org/wiki/XXTEA 

**Fully compatible** to use for encryption / decryption evaluation of data of the**  [ boseji xxtea Cryptography library for Arduino](https://github.com/boseji/xxtea-lib "xxtea Cryptography library for Arduino")** on Windows PC

**WARNING! **Do not use outdated implementation https://www.tools4noobs.com/online_tools/xxtea_encrypt/ to evaluate this tool

------------
# Usage

### String encryption 
##### Example
```bash
xxtea-tool -input=SomeRandomStringToEncrypt -key=AnyKeyYouLike123
```

To encrypt something, provide **message (string)** in the `-input ` and **key** in the  `-key`  arguments. Order does not matter.

If you want just to play arround you may omit providing key, and everything will be encrypted with the default `0123456789abcdef` key. 

Specify additional `-s` option for silent operation.

Encryption key should not be longer then 16 characters (represeting 16 bytes). Leading and trailing spaces are trimmed. 

In case key is smaller than 16 bytes, trailing zeroes are appended automatically to match 16 byte key size.

------------



### String decryption 
##### Example
```bash
xxtea-tool -d -input=EncryptedMessage -key=DecryptionKey123
```
**Note: Argument `-d` is mandatory for decryption**, as by default tool performs encryption operation.

Other than that,  as with the encryption, providing  `-input ` and  `-key` arguments is neccessary (unless it was encrypted with default `0123456789abcdef` key). Adding additional `-s` flag runs tool in a silent mode.

------------



### Other arguments

`-s` : Silent mode, all information messages will be suspended. Tool will output only a result of conversion (encryption or decryption). Usefull for batch operations or chaining with external executable program for the pipeline.  Used as an additional option within normal encryption / decryption operation.

`-v` : Prints program version

`-h` : Prints help screen with all of the available options

### License
BSD 2-Clause "Simplified" License

### Credits
Corrected Block TEA / XXTEA library in Go by [hillu](https://github.com/hillu "hillu")
https://github.com/hillu/go-xxtea


