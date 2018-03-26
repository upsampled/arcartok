

# ascart
`import "github.com/liamCDI/ascartok"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>
Package ascart Package Ascii Art Overkill converts an image
to ascii art for a given `pallet` of ascii characters.

This is done first by running the image through the golang image transformation library `gift`
The Filter Options string follows the following format `<gift filter name>=<arg1>.<arg2>;<gift filter name>=<arg1>`
The Filters are run over the image and then the outputed image is sent to the ascii converter

The ascii converter takes the RGB values from the image and converts them to Luma. The Luma is then
linearly mapped to the given pallet string (IE: luma=0 maps to the first character in teh pallet string)

The command `ascartok` exposes the main method of the library and libascart is a c library wrapper.

Example usage: ascartok -in ./test/gopher.png -plt=" .:OI M" -flt="resize=50,0"

 
```go
img=`			
                              :II        OOOO O:I .:I::I .      OI O.
                          ..I      :O:OIOIIOOOOI   IIIOOOO::III IIII IOOO:
                       :       I O:II   IOOI: MMMMMMMMMM MOI IO:M ..O M  OI O
                   .       :I II   OIO.MMIMMMMMMMMMMMMIOOOM OOOMO..  : MOI.
                .O     ::III   OI.M    : MMMMMMMMMMMM    .:O:I     I::IMMI ..
              :I    :IOOIII:O:M       I:MMMMMMMMMMMMMO.  .:MOO       IOO:
             O   : OOOI:MOM IIIIIII    :OMMMMMMMMMMMMMMMMMMM:I          :    .
               O I:O .MMMMMMMMMMM II    :IMMMMMMMMMMMMMMMM :I             OO  .
        II:  :O IM:MMMMMMMMMMI:..:O:O    IOI MMMMMMMMMMM OO              .OI:
     I      .I.IMMMMMMMMMMMM     :  :O     IOOOOIIIIIOIII                  O I
        :I: II I.MMMMMMMMMMMMI::::IMI: O:.....OIOOOI                        .M.  .
     O :I I.M   : MMMMMMMMMMMMMMMMMM:O       :OOOOO:O                        ..:  :
   :.O  MOIM    I: MMMMMMMMMMMMMMM OOO.::::OIIIIIII:.                         .IO  I
  I :.OM:  O     IOOI MMMMMMMMM IOOI:OIIIIIOO:OIO.:OI                          .IO
  O O II IM:        IIOOIIIIIIIII   ::OOO:OI IIMM OI                            . OI
    O:II  IM                         OOOIO:MMM IMMM:I                            ..II   .    M
      ::O.:.I                             IOMMMO:IIII                              O .
           O:OM                            IOIIIII                                  MO.I:MO. O:
       II   :IOM                                                                      M:MMMMM :I .
           : .O.M                                                                    I ::: IMOI O
               OMM                                                                     O
             : :OMI                                                                    M:OI   O
              I   IO.                                           IOIOIIII                  I::
               I  OI O                                          :MMMMM  OI              OM  .
                O   .I:M                                     IIOOM  MMMI:I              .: .
                     M:I:                                   I:OOIIIOOIIOO                I I
                        O:                                  I        II                 MI   .
                        IO.                                                             MII
                         IO.                                                            .II
                       .  :M.                                                           MIM  .
                        .  .I                                                           MI.  .
                            MM                                                          MIO.
                         . : I                                                          :M
                            .IM                                                         :.O .
                          ..:O.                                                        MM:I
                          :  IO                                                       I::: :
                          : OIO                                                       MMM
                          . .OO                                                      MO:
                           ::O.                                                    .O :IMM OO
                          : .M.                                                 MO. I:OMMMMM .  .
                          : OO.I                                            M.I. IOOO  OI  .
                             .OOO                                       M.IO  II O
                          : IMIO .OM                               M.II.  II : O    O . :
                          .  :::II  IOM                     MM.OIOM  III:M O     I:       OI
                                 I OI          M.....OOIIIOM   IIIO.II.:      IO
                            M:II    I: IOIIOO:I        IIIIOOI:::O        O .
                                 O      I MM  :OI: OII:::I           O:O
                                   :: :.MMMMM:.                :. :
                                  .. .MIMMMIO         OOIII:
                                  .  O:.: .O
`
```
 

## <a name="pkg-index">Index</a>
* [type AscArt](#AscArt)
  * [func Img2asc(fn string, plte string, flarg string) (*AscArt, error)](#Img2asc)
  * [func (a *AscArt) String() string](#AscArt.String)


#### <a name="pkg-files">Package files</a>
[giftwrapper.go](/src/github.com/liamCDI/ascartok/giftwrapper.go) [image2pix.go](/src/github.com/liamCDI/ascartok/image2pix.go) 






## <a name="AscArt">type</a> [AscArt](/src/target/image2pix.go?s=3771:3852#L87)
``` go
type AscArt struct {
    H   int    //Height
    W   int    //Width
    Art string //Art
}
```
AscArt holds the ascii art







### <a name="Img2asc">func</a> [Img2asc](/src/target/image2pix.go?s=4100:4167#L106)
``` go
func Img2asc(fn string, plte string, flarg string) (*AscArt, error)
```
Img2asc returns an ascii art object for a file name





### <a name="AscArt.String">func</a> (\*AscArt) [String](/src/target/image2pix.go?s=3854:3886#L93)
``` go
func (a *AscArt) String() string
```







- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
