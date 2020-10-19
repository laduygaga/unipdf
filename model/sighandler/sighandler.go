//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

// Package sighandler implements digital signature handlers for PDF signature validation and signing.
package sighandler ;import (_f "bytes";_gd "crypto";_ge "crypto/rand";_fd "crypto/rsa";_ddb "crypto/x509";_de "crypto/x509/pkix";_b "encoding/asn1";_ddf "errors";_gdg "fmt";_gg "github.com/unidoc/pkcs7";_bg "github.com/unidoc/timestamp";_fce "github.com/unidoc/unipdf/v3/core";_e "github.com/unidoc/unipdf/v3/model";_dd "hash";_a "io";_c "io/ioutil";_g "net/http";_fc "time";);

// NewAdobePKCS7Detached creates a new Adobe.PPKMS/Adobe.PPKLite adbe.pkcs7.detached signature handler.
// Both parameters may be nil for the signature validation.
func NewAdobePKCS7Detached (privateKey *_fd .PrivateKey ,certificate *_ddb .Certificate )(_e .SignatureHandler ,error ){return &adobePKCS7Detached {_fg :certificate ,_fe :privateKey },nil ;};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_ebb *adobeX509RSASHA1 )IsApplicable (sig *_e .PdfSignature )bool {if sig ==nil ||sig .Filter ==nil ||sig .SubFilter ==nil {return false ;};return (*sig .Filter =="A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053"||*sig .Filter =="\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")&&*sig .SubFilter =="\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031";};func (_dc *adobePKCS7Detached )getCertificate (_cf *_e .PdfSignature )(*_ddb .Certificate ,error ){if _dc ._fg !=nil {return _dc ._fg ,nil ;};var _fa []byte ;switch _bb :=_cf .Cert .(type ){case *_fce .PdfObjectString :_fa =_bb .Bytes ();case *_fce .PdfObjectArray :if _bb .Len ()==0{return nil ,_ddf .New ("\u006e\u006f\u0020s\u0069\u0067\u006e\u0061t\u0075\u0072\u0065\u0020\u0063\u0065\u0072t\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0073\u0020\u0066\u006f\u0075\u006e\u0064");};for _ ,_be :=range _bb .Elements (){_bde ,_bf :=_fce .GetString (_be );if !_bf {return nil ,_gdg .Errorf ("\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0074\u0079p\u0065\u0020\u0069\u006e\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065 \u0063\u0065r\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u0063h\u0061\u0069\u006e\u003a\u0020\u0025\u0054",_be );};_fa =append (_fa ,_bde .Bytes ()...);};default:return nil ,_gdg .Errorf ("\u0069n\u0076\u0061l\u0069\u0064\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072e\u0020\u0063\u0065\u0072\u0074\u0069f\u0069\u0063\u0061\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054",_bb );};_ggb ,_fb :=_ddb .ParseCertificates (_fa );if _fb !=nil {return nil ,_fb ;};return _ggb [0],nil ;};

// Sign sets the Contents fields.
func (_ef *adobePKCS7Detached )Sign (sig *_e .PdfSignature ,digest _e .Hasher )error {if _ef ._gdgd {_dcc :=_ef ._fdc ;if _dcc <=0{_dcc =8192;};sig .Contents =_fce .MakeHexString (string (make ([]byte ,_dcc )));return nil ;};_ded :=digest .(*_f .Buffer );_ab ,_ggbd :=_gg .NewSignedData (_ded .Bytes ());if _ggbd !=nil {return _ggbd ;};if _ae :=_ab .AddSigner (_ef ._fg ,_ef ._fe ,_gg .SignerInfoConfig {});_ae !=nil {return _ae ;};_ab .Detach ();_fed ,_ggbd :=_ab .Finish ();if _ggbd !=nil {return _ggbd ;};_bdf :=make ([]byte ,8192);copy (_bdf ,_fed );sig .Contents =_fce .MakeHexString (string (_bdf ));return nil ;};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_efe *docTimeStamp )IsApplicable (sig *_e .PdfSignature )bool {if sig ==nil ||sig .Filter ==nil ||sig .SubFilter ==nil {return false ;};return (*sig .Filter =="A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053"||*sig .Filter =="\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")&&*sig .SubFilter =="\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031";};

// NewDigest creates a new digest.
func (_eb *adobePKCS7Detached )NewDigest (sig *_e .PdfSignature )(_e .Hasher ,error ){return _f .NewBuffer (nil ),nil ;};

// InitSignature initialises the PdfSignature.
func (_gacg *docTimeStamp )InitSignature (sig *_e .PdfSignature )error {_dfc :=*_gacg ;sig .Handler =&_dfc ;sig .Filter =_fce .MakeName ("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065");sig .SubFilter =_fce .MakeName ("\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031");sig .Reference =nil ;if _gacg ._da > 0{sig .Contents =_fce .MakeHexString (string (make ([]byte ,_gacg ._da )));}else {_bbb ,_fae :=_gacg .NewDigest (sig );if _fae !=nil {return _fae ;};_bbb .Write ([]byte ("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"));if _fae =_dfc .Sign (sig ,_bbb );_fae !=nil {return _fae ;};_gacg ._da =_dfc ._da ;};return nil ;};type adobeX509RSASHA1 struct{_gf *_fd .PrivateKey ;_gdgdb *_ddb .Certificate ;_fad SignFunc ;};

// DocTimeStampOpts defines options for configuring the timestamp handler.
type DocTimeStampOpts struct{

// SignatureSize is the estimated size of the signature contents in bytes.
// If not provided, a default signature size of 4192 is used.
// The signing process will report the model.ErrSignNotEnoughSpace error
// if the estimated signature size is smaller than the actual size of the
// signature.
SignatureSize int ;};

// Validate validates PdfSignature.
func (_gee *docTimeStamp )Validate (sig *_e .PdfSignature ,digest _e .Hasher )(_e .SignatureValidationResult ,error ){_dad :=sig .Contents .Bytes ();_ebdbb ,_def :=_gg .Parse (_dad );if _def !=nil {return _e .SignatureValidationResult {},_def ;};if _def =_ebdbb .Verify ();_def !=nil {return _e .SignatureValidationResult {},_def ;};var _cb timestampInfo ;_ ,_def =_b .Unmarshal (_ebdbb .Content ,&_cb );if _def !=nil {return _e .SignatureValidationResult {},_def ;};_ebf ,_def :=_db (_cb .MessageImprint .HashAlgorithm .Algorithm );if _def !=nil {return _e .SignatureValidationResult {},_def ;};_cca :=_ebf .New ();_dgd :=digest .(*_f .Buffer );_cca .Write (_dgd .Bytes ());_ccf :=_cca .Sum (nil );_bge :=_e .SignatureValidationResult {IsSigned :true ,IsVerified :_f .Equal (_ccf ,_cb .MessageImprint .HashedMessage ),GeneralizedTime :_cb .GeneralizedTime };return _bge ,nil ;};func _db (_aed _b .ObjectIdentifier )(_gd .Hash ,error ){switch {case _aed .Equal (_gg .OIDDigestAlgorithmSHA1 ),_aed .Equal (_gg .OIDDigestAlgorithmECDSASHA1 ),_aed .Equal (_gg .OIDDigestAlgorithmDSA ),_aed .Equal (_gg .OIDDigestAlgorithmDSASHA1 ),_aed .Equal (_gg .OIDEncryptionAlgorithmRSA ):return _gd .SHA1 ,nil ;case _aed .Equal (_gg .OIDDigestAlgorithmSHA256 ),_aed .Equal (_gg .OIDDigestAlgorithmECDSASHA256 ):return _gd .SHA256 ,nil ;case _aed .Equal (_gg .OIDDigestAlgorithmSHA384 ),_aed .Equal (_gg .OIDDigestAlgorithmECDSASHA384 ):return _gd .SHA384 ,nil ;case _aed .Equal (_gg .OIDDigestAlgorithmSHA512 ),_aed .Equal (_gg .OIDDigestAlgorithmECDSASHA512 ):return _gd .SHA512 ,nil ;};return _gd .Hash (0),_gg .ErrUnsupportedAlgorithm ;};type timestampInfo struct{Version int ;Policy _b .RawValue ;MessageImprint struct{HashAlgorithm _de .AlgorithmIdentifier ;HashedMessage []byte ;};SerialNumber _b .RawValue ;GeneralizedTime _fc .Time ;};

// NewDigest creates a new digest.
func (_bee *docTimeStamp )NewDigest (sig *_e .PdfSignature )(_e .Hasher ,error ){return _f .NewBuffer (nil ),nil ;};

// NewAdobeX509RSASHA1Custom creates a new Adobe.PPKMS/Adobe.PPKLite adbe.x509.rsa_sha1 signature handler
// with a custom signing function. Both parameters may be nil for the signature validation.
func NewAdobeX509RSASHA1Custom (certificate *_ddb .Certificate ,signFunc SignFunc )(_e .SignatureHandler ,error ){return &adobeX509RSASHA1 {_gdgdb :certificate ,_fad :signFunc },nil ;};

// SignFunc represents a custom signing function. The function should return
// the computed signature.
type SignFunc func (_dfg *_e .PdfSignature ,_ddc _e .Hasher )([]byte ,error );

// InitSignature initialises the PdfSignature.
func (_fee *adobeX509RSASHA1 )InitSignature (sig *_e .PdfSignature )error {if _fee ._gdgdb ==nil {return _ddf .New ("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c");};if _fee ._gf ==nil &&_fee ._fad ==nil {return _ddf .New ("\u006d\u0075\u0073\u0074\u0020\u0070\u0072o\u0076\u0069\u0064e\u0020\u0065\u0069t\u0068\u0065r\u0020\u0061\u0020\u0070\u0072\u0069v\u0061te\u0020\u006b\u0065\u0079\u0020\u006f\u0072\u0020\u0061\u0020\u0073\u0069\u0067\u006e\u0069\u006e\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e");};_ebd :=*_fee ;sig .Handler =&_ebd ;sig .Filter =_fce .MakeName ("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065");sig .SubFilter =_fce .MakeName ("\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031");sig .Cert =_fce .MakeString (string (_ebd ._gdgdb .Raw ));sig .Reference =nil ;_ga ,_faa :=_ebd .NewDigest (sig );if _faa !=nil {return _faa ;};_ga .Write ([]byte ("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"));return _ebd .Sign (sig ,_ga );};

// Sign sets the Contents fields for the PdfSignature.
func (_acg *docTimeStamp )Sign (sig *_e .PdfSignature ,digest _e .Hasher )error {_adf :=digest .(*_f .Buffer );_gacd :=_acg ._ffg .New ();if _ ,_fde :=_a .Copy (_gacd ,_adf );_fde !=nil {return _fde ;};_cdc :=_bg .Request {HashAlgorithm :_acg ._ffg ,HashedMessage :_gacd .Sum (nil ),Certificates :true ,Extensions :nil ,ExtraExtensions :nil };_ggfg ,_feg :=_cdc .Marshal ();if _feg !=nil {return _feg ;};_agg ,_feg :=_g .Post (_acg ._cd ,"a\u0070\u0070\u006c\u0069\u0063\u0061t\u0069\u006f\u006e\u002f\u0074\u0069\u006d\u0065\u0073t\u0061\u006d\u0070-\u0071u\u0065\u0072\u0079",_f .NewBuffer (_ggfg ));if _feg !=nil {return _feg ;};defer _agg .Body .Close ();_ffae ,_feg :=_c .ReadAll (_agg .Body );if _feg !=nil {return _feg ;};if _agg .StatusCode !=_g .StatusOK {return _gdg .Errorf ("\u0068\u0074\u0074\u0070\u0020\u0073\u0074\u0061\u0074\u0075\u0073\u0020\u0063\u006f\u0064e\u0020n\u006f\u0074\u0020\u006f\u006b\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0064\u0029",_agg .StatusCode );};var _gff struct{Version _b .RawValue ;Content _b .RawValue ;};if _ ,_feg =_b .Unmarshal (_ffae ,&_gff );_feg !=nil {return _feg ;};_afc :=len (_gff .Content .FullBytes );if _acg ._da > 0&&_afc > _acg ._da {return _e .ErrSignNotEnoughSpace ;};if _afc > 0{_acg ._da =_afc +128;};sig .Contents =_fce .MakeHexString (string (_gff .Content .FullBytes ));return nil ;};func (_aa *adobeX509RSASHA1 )getCertificate (_gac *_e .PdfSignature )(*_ddb .Certificate ,error ){if _aa ._gdgdb !=nil {return _aa ._gdgdb ,nil ;};var _gfg []byte ;switch _cfc :=_gac .Cert .(type ){case *_fce .PdfObjectString :_gfg =_cfc .Bytes ();case *_fce .PdfObjectArray :if _cfc .Len ()==0{return nil ,_ddf .New ("\u006e\u006f\u0020s\u0069\u0067\u006e\u0061t\u0075\u0072\u0065\u0020\u0063\u0065\u0072t\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0073\u0020\u0066\u006f\u0075\u006e\u0064");};for _ ,_efb :=range _cfc .Elements (){_gace ,_gec :=_fce .GetString (_efb );if !_gec {return nil ,_gdg .Errorf ("\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0074\u0079p\u0065\u0020\u0069\u006e\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065 \u0063\u0065r\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u0063h\u0061\u0069\u006e\u003a\u0020\u0025\u0054",_efb );};_gfg =append (_gfg ,_gace .Bytes ()...);};default:return nil ,_gdg .Errorf ("\u0069n\u0076\u0061l\u0069\u0064\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072e\u0020\u0063\u0065\u0072\u0074\u0069f\u0069\u0063\u0061\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054",_cfc );};_ad ,_ce :=_ddb .ParseCertificates (_gfg );if _ce !=nil {return nil ,_ce ;};return _ad [0],nil ;};

// NewAdobeX509RSASHA1 creates a new Adobe.PPKMS/Adobe.PPKLite adbe.x509.rsa_sha1 signature handler.
// Both parameters may be nil for the signature validation.
func NewAdobeX509RSASHA1 (privateKey *_fd .PrivateKey ,certificate *_ddb .Certificate )(_e .SignatureHandler ,error ){return &adobeX509RSASHA1 {_gdgdb :certificate ,_gf :privateKey },nil ;};

// Sign sets the Contents fields for the PdfSignature.
func (_ggf *adobeX509RSASHA1 )Sign (sig *_e .PdfSignature ,digest _e .Hasher )error {var _ebdb []byte ;var _ba error ;if _ggf ._fad !=nil {_ebdb ,_ba =_ggf ._fad (sig ,digest );if _ba !=nil {return _ba ;};}else {_fbe ,_afa :=digest .(_dd .Hash );if !_afa {return _ddf .New ("\u0068a\u0073h\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");};_dg ,_ :=_aca (_ggf ._gdgdb .SignatureAlgorithm );_ebdb ,_ba =_fd .SignPKCS1v15 (_ge .Reader ,_ggf ._gf ,_dg ,_fbe .Sum (nil ));if _ba !=nil {return _ba ;};};_ebdb ,_ba =_b .Marshal (_ebdb );if _ba !=nil {return _ba ;};sig .Contents =_fce .MakeHexString (string (_ebdb ));return nil ;};type docTimeStamp struct{_cd string ;_ffg _gd .Hash ;_da int ;};

// NewDocTimeStamp creates a new DocTimeStamp signature handler.
// Both the timestamp server URL and the hash algorithm can be empty for the
// signature validation.
// The following hash algorithms are supported:
// crypto.SHA1, crypto.SHA256, crypto.SHA384, crypto.SHA512.
// NOTE: the handler will do a mock Sign when initializing the signature
// in order to estimate the signature size. Use NewDocTimeStampWithOpts
// for providing the signature size.
func NewDocTimeStamp (timestampServerURL string ,hashAlgorithm _gd .Hash )(_e .SignatureHandler ,error ){return &docTimeStamp {_cd :timestampServerURL ,_ffg :hashAlgorithm },nil ;};

// Validate validates PdfSignature.
func (_ff *adobePKCS7Detached )Validate (sig *_e .PdfSignature ,digest _e .Hasher )(_e .SignatureValidationResult ,error ){_df :=sig .Contents .Bytes ();_ea ,_ee :=_gg .Parse (_df );if _ee !=nil {return _e .SignatureValidationResult {},_ee ;};_bdg :=digest .(*_f .Buffer );_ea .Content =_bdg .Bytes ();if _ee =_ea .Verify ();_ee !=nil {return _e .SignatureValidationResult {},_ee ;};return _e .SignatureValidationResult {IsSigned :true ,IsVerified :true },nil ;};

// NewEmptyAdobePKCS7Detached creates a new Adobe.PPKMS/Adobe.PPKLite adbe.pkcs7.detached
// signature handler. The generated signature is empty and of size signatureLen.
// The signatureLen parameter can be 0 for the signature validation.
func NewEmptyAdobePKCS7Detached (signatureLen int )(_e .SignatureHandler ,error ){return &adobePKCS7Detached {_gdgd :true ,_fdc :signatureLen },nil ;};

// NewDocTimeStampWithOpts returns a new DocTimeStamp configured using the
// specified options. If no options are provided, default options will be used.
// Both the timestamp server URL and the hash algorithm can be empty for the
// signature validation.
// The following hash algorithms are supported:
// crypto.SHA1, crypto.SHA256, crypto.SHA384, crypto.SHA512.
func NewDocTimeStampWithOpts (timestampServerURL string ,hashAlgorithm _gd .Hash ,opts *DocTimeStampOpts )(_e .SignatureHandler ,error ){if opts ==nil {opts =&DocTimeStampOpts {};};if opts .SignatureSize <=0{opts .SignatureSize =4192;};return &docTimeStamp {_cd :timestampServerURL ,_ffg :hashAlgorithm ,_da :opts .SignatureSize },nil ;};

// InitSignature initialises the PdfSignature.
func (_bd *adobePKCS7Detached )InitSignature (sig *_e .PdfSignature )error {if !_bd ._gdgd {if _bd ._fg ==nil {return _ddf .New ("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c");};if _bd ._fe ==nil {return _ddf .New ("\u0070\u0072\u0069\u0076\u0061\u0074\u0065\u004b\u0065\u0079\u0020m\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0062\u0065 \u006e\u0069\u006c");};};_cg :=*_bd ;sig .Handler =&_cg ;sig .Filter =_fce .MakeName ("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065");sig .SubFilter =_fce .MakeName ("\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064");sig .Reference =nil ;_ac ,_cge :=_cg .NewDigest (sig );if _cge !=nil {return _cge ;};_ac .Write ([]byte ("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"));return _cg .Sign (sig ,_ac );};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature
func (_dcg *adobePKCS7Detached )IsApplicable (sig *_e .PdfSignature )bool {if sig ==nil ||sig .Filter ==nil ||sig .SubFilter ==nil {return false ;};return (*sig .Filter =="A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053"||*sig .Filter =="\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")&&*sig .SubFilter =="\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064";};

// Validate validates PdfSignature.
func (_cc *adobeX509RSASHA1 )Validate (sig *_e .PdfSignature ,digest _e .Hasher )(_e .SignatureValidationResult ,error ){_af ,_eg :=_cc .getCertificate (sig );if _eg !=nil {return _e .SignatureValidationResult {},_eg ;};_fdcb :=sig .Contents .Bytes ();var _fgg []byte ;if _ ,_ffa :=_b .Unmarshal (_fdcb ,&_fgg );_ffa !=nil {return _e .SignatureValidationResult {},_ffa ;};_bca ,_fcg :=digest .(_dd .Hash );if !_fcg {return _e .SignatureValidationResult {},_ddf .New ("\u0068a\u0073h\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");};_eaa ,_ :=_aca (_af .SignatureAlgorithm );if _gc :=_fd .VerifyPKCS1v15 (_af .PublicKey .(*_fd .PublicKey ),_eaa ,_bca .Sum (nil ),_fgg );_gc !=nil {return _e .SignatureValidationResult {},_gc ;};return _e .SignatureValidationResult {IsSigned :true ,IsVerified :true },nil ;};type adobePKCS7Detached struct{_fe *_fd .PrivateKey ;_fg *_ddb .Certificate ;_gdgd bool ;_fdc int ;};func _aca (_beb _ddb .SignatureAlgorithm )(_gd .Hash ,bool ){var _ca _gd .Hash ;switch _beb {case _ddb .SHA1WithRSA :_ca =_gd .SHA1 ;case _ddb .SHA256WithRSA :_ca =_gd .SHA256 ;case _ddb .SHA384WithRSA :_ca =_gd .SHA384 ;case _ddb .SHA512WithRSA :_ca =_gd .SHA512 ;default:return _gd .SHA1 ,false ;};return _ca ,true ;};

// NewDigest creates a new digest.
func (_aad *adobeX509RSASHA1 )NewDigest (sig *_e .PdfSignature )(_e .Hasher ,error ){_gab ,_gdf :=_aad .getCertificate (sig );if _gdf !=nil {return nil ,_gdf ;};_fff ,_ :=_aca (_gab .SignatureAlgorithm );return _fff .New (),nil ;};func (_daa *docTimeStamp )getCertificate (_ag *_e .PdfSignature )(*_ddb .Certificate ,error ){var _dfa []byte ;switch _bfg :=_ag .Cert .(type ){case *_fce .PdfObjectString :_dfa =_bfg .Bytes ();case *_fce .PdfObjectArray :if _bfg .Len ()==0{return nil ,_ddf .New ("\u006e\u006f\u0020s\u0069\u0067\u006e\u0061t\u0075\u0072\u0065\u0020\u0063\u0065\u0072t\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0073\u0020\u0066\u006f\u0075\u006e\u0064");};for _ ,_bba :=range _bfg .Elements (){_agb ,_fab :=_fce .GetString (_bba );if !_fab {return nil ,_gdg .Errorf ("\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0074\u0079p\u0065\u0020\u0069\u006e\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065 \u0063\u0065r\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u0063h\u0061\u0069\u006e\u003a\u0020\u0025\u0054",_bba );};_dfa =append (_dfa ,_agb .Bytes ()...);};default:return nil ,_gdg .Errorf ("\u0069n\u0076\u0061l\u0069\u0064\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072e\u0020\u0063\u0065\u0072\u0074\u0069f\u0069\u0063\u0061\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054",_bfg );};_afe ,_ddcb :=_ddb .ParseCertificates (_dfa );if _ddcb !=nil {return nil ,_ddcb ;};return _afe [0],nil ;};