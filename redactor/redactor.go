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

package redactor ;import (_a "errors";_e "fmt";_ee "github.com/unidoc/unipdf/v3/common";_d "github.com/unidoc/unipdf/v3/contentstream";_fe "github.com/unidoc/unipdf/v3/core";_gd "github.com/unidoc/unipdf/v3/creator";_f "github.com/unidoc/unipdf/v3/extractor";
_c "github.com/unidoc/unipdf/v3/model";_ag "io";_ga "regexp";_b "sort";_gc "strings";);func _edbga (_cga []placeHolders )[]replacement {_eeaa :=[]replacement {};for _ ,_cbe :=range _cga {_agd :=_cbe ._db ;_adab :=_cbe ._cc ;_ffd :=_cbe ._dc ;for _ ,_gfa :=range _agd {_gbd :=replacement {_ec :_adab ,_gag :_ffd ,_ed :_gfa };
_eeaa =append (_eeaa ,_gbd );};};_b .Slice (_eeaa ,func (_afe ,_ffaa int )bool {return _eeaa [_afe ]._ed < _eeaa [_ffaa ]._ed });return _eeaa ;};func _aec (_cgd localSpanMarks ,_afd *_f .TextMarkArray ,_abd *_c .PdfFont ,_fgc ,_edbg string )([]_fe .PdfObject ,error ){_fdc :=_fbg (_afd );
Tj ,_cad :=_bfe (_afd );if _cad !=nil {return nil ,_cad ;};_ffa :=len (_fgc );_ged :=len (_fdc );_bad :=-1;_fde :=_fe .MakeFloat (Tj );if _fdc !=_edbg {_deb :=_cgd ._afdd ;if _deb ==0{_bad =_gc .LastIndex (_fgc ,_fdc );}else {_bad =_gc .Index (_fgc ,_fdc );
};}else {_bad =_gc .Index (_fgc ,_fdc );};_cee :=_bad +_ged ;_eab :=[]_fe .PdfObject {};if _bad ==0&&_cee ==_ffa {_eab =append (_eab ,_fde );}else if _bad ==0&&_cee < _ffa {_ceec :=_ca (_fgc [_cee :],_abd );_daf :=_fe .MakeStringFromBytes (_ceec );_eab =append (_eab ,_fde ,_daf );
}else if _bad > 0&&_cee >=_ffa {_fbc :=_ca (_fgc [:_bad ],_abd );_fgcf :=_fe .MakeStringFromBytes (_fbc );_eab =append (_eab ,_fgcf ,_fde );}else if _bad > 0&&_cee < _ffa {_eddb :=_ca (_fgc [:_bad ],_abd );_bgf :=_ca (_fgc [_cee :],_abd );_ebdc :=_fe .MakeStringFromBytes (_eddb );
_dcg :=_fe .MakeString (string (_bgf ));_eab =append (_eab ,_ebdc ,_fde ,_dcg );};return _eab ,nil ;};func _ecdc (_fdd ,_ggd ,_eaed float64 )float64 {_eaed =_eaed /100;_bag :=(-1000*_fdd )/(_ggd *_eaed );return _bag ;};func _bf (_gf *_d .ContentStreamOperations ,_gbe string ,_dcf int )error {_eb :=_d .ContentStreamOperations {};
var _gdc _d .ContentStreamOperation ;for _df ,_cg :=range *_gf {if _df ==_dcf {if _gbe =="\u0027"{_cd :=_d .ContentStreamOperation {Operand :"\u0054\u002a"};_eb =append (_eb ,&_cd );_gdc .Params =_cg .Params ;_gdc .Operand ="\u0054\u006a";_eb =append (_eb ,&_gdc );
}else if _gbe =="\u0022"{_ad :=_cg .Params [:2];Tc ,Tw :=_ad [0],_ad [1];_ebd :=_d .ContentStreamOperation {Params :[]_fe .PdfObject {Tc },Operand :"\u0054\u0063"};_eb =append (_eb ,&_ebd );_ebd =_d .ContentStreamOperation {Params :[]_fe .PdfObject {Tw },Operand :"\u0054\u0077"};
_eb =append (_eb ,&_ebd );_gdc .Params =[]_fe .PdfObject {_cg .Params [2]};_gdc .Operand ="\u0054\u006a";_eb =append (_eb ,&_gdc );};};_eb =append (_eb ,_cg );};*_gf =_eb ;return nil ;};func _fdcc (_aaad string ,_cbg []localSpanMarks )([]placeHolders ,error ){_dbega :="";
_ade :=[]placeHolders {};for _aef ,_gcg :=range _cbg {_edfe :=_gcg ._gbde ;_bcb :=_gcg ._cfbc ;_bgd :=_fbg (_edfe );_geb ,_gbfc :=_bfe (_edfe );if _gbfc !=nil {return nil ,_gbfc ;};if _bgd !=_dbega {var _eee []int ;if _aef ==0&&_bcb !=_bgd {_dfcg :=_gc .Index (_aaad ,_bgd );
_eee =[]int {_dfcg };}else if _aef ==len (_cbg )-1{_faea :=_gc .LastIndex (_aaad ,_bgd );_eee =[]int {_faea };}else {_eee =_ebbg (_aaad ,_bgd );};_dafa :=placeHolders {_db :_eee ,_cc :_bgd ,_dc :_geb };_ade =append (_ade ,_dafa );};_dbega =_bgd ;};return _ade ,nil ;
};func _ecdd (_gdea *_d .ContentStreamOperations ,PdfObj _fe .PdfObject )(*_d .ContentStreamOperation ,int ,bool ){for _ecgc ,_eff :=range *_gdea {_cae :=_eff .Operand ;if _cae =="\u0054\u006a"{_fee :=_fe .TraceToDirectObject (_eff .Params [0]);if _fee ==PdfObj {return _eff ,_ecgc ,true ;
};}else if _cae =="\u0054\u004a"{_def ,_bcf :=_fe .GetArray (_eff .Params [0]);if !_bcf {return nil ,_ecgc ,_bcf ;};for _ ,_cbcd :=range _def .Elements (){if _cbcd ==PdfObj {return _eff ,_ecgc ,true ;};};}else if _cae =="\u0022"{_dcegc :=_fe .TraceToDirectObject (_eff .Params [2]);
if _dcegc ==PdfObj {return _eff ,_ecgc ,true ;};}else if _cae =="\u0027"{_fbef :=_fe .TraceToDirectObject (_eff .Params [0]);if _fbef ==PdfObj {return _eff ,_ecgc ,true ;};};};return nil ,-1,false ;};func _ccc (_fbf string ,_aeg []replacement ,_gdcd *_c .PdfFont )[]_fe .PdfObject {_dcad :=[]_fe .PdfObject {};
_fcf :=0;_ef :=_fbf ;for _cbc ,_gee :=range _aeg {_dgec :=_gee ._ed ;_bda :=_gee ._gag ;_bbc :=_gee ._ec ;_aegg :=_fe .MakeFloat (_bda );_gba :=_fbf [_fcf :_dgec ];_adcg :=_ca (_gba ,_gdcd );_edc :=_fe .MakeStringFromBytes (_adcg );_dcad =append (_dcad ,_edc );
_dcad =append (_dcad ,_aegg );_ebb :=_dgec +len (_bbc );_ef =_fbf [_ebb :];_fcf =_ebb ;if _cbc ==len (_aeg )-1{_adcg =_ca (_ef ,_gdcd );_edc =_fe .MakeStringFromBytes (_adcg );_dcad =append (_dcad ,_edc );};};return _dcad ;};func _gdcb (_feff *_f .TextMarkArray )*_c .PdfFont {_ ,_dggd :=_dce (_feff );
_gcf :=_feff .Elements ()[_dggd ];_dfc :=_gcf .Font ;return _dfc ;};

// RedactionOptions is a collection of RedactionTerm objects.
type RedactionOptions struct{Terms []RedactionTerm ;};type localSpanMarks struct{_gbde *_f .TextMarkArray ;_afdd int ;_cfbc string ;};type matchedIndex struct{_cec int ;_cagg int ;_fdg string ;};func _cbd (_eea *_c .PdfFont ,_cde _f .TextMark )float64 {_gg :=0.001;
_fab :=_cde .Th /100;if _eea .Subtype ()=="\u0054\u0079\u0070e\u0033"{_gg =1;};_ab ,_gfd :=_eea .GetRuneMetrics (' ');if !_gfd {_ab ,_gfd =_eea .GetCharMetrics (32);};if !_gfd {_ab ,_ =_c .DefaultFont ().GetRuneMetrics (' ');};_afac :=_gg *((_ab .Wx *_cde .FontSize +_cde .Tc +_cde .Tw )/_fab );
return _afac ;};

// Redact executes the redact operation on a pdf file and updates the content streams of all pages of the file.
func (_gaa *Redactor )Redact ()error {_gbec ,_dfcf :=_gaa ._egf .GetNumPages ();if _dfcf !=nil {return _e .Errorf ("\u0066\u0061\u0069\u006c\u0065\u0064 \u0074\u006f\u0020\u0067\u0065\u0074\u0020\u0074\u0068\u0065\u0020\u006e\u0075m\u0062\u0065\u0072\u0020\u006f\u0066\u0020P\u0061\u0067\u0065\u0073");
};_bga :=_gaa ._cbcf .FillColor ;_fbca :=_gaa ._cbcf .BorderWidth ;_egbf :=_gaa ._cbcf .FillOpacity ;for _ggfc :=1;_ggfc <=_gbec ;_ggfc ++{_adcd ,_ebda :=_gaa ._egf .GetPage (_ggfc );if _ebda !=nil {return _ebda ;};_gda ,_ebda :=_f .New (_adcd );if _ebda !=nil {return _ebda ;
};_eaf ,_ ,_ ,_ebda :=_gda .ExtractPageText ();if _ebda !=nil {return _ebda ;};_cab :=_eaf .GetContentStreamOps ();_gga ,_gaab ,_ebda :=_gaa .redactPage (_cab ,_adcd .Resources );if _gaab ==nil {_ee .Log .Info ("N\u006f\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0066\u006f\u0072\u0020t\u0068\u0065\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065d \u0070\u0061\u0074t\u0061r\u006e\u002e");
_gaab =_cab ;};_adcd .SetContentStreams ([]string {_gaab .String ()},_fe .NewFlateEncoder ());if _ebda !=nil {return _ebda ;};_aefg ,_ebda :=_adcd .GetMediaBox ();if _ebda !=nil {return _ebda ;};if _adcd .MediaBox ==nil {_adcd .MediaBox =_aefg ;};if _faed :=_gaa ._egcf .AddPage (_adcd );
_faed !=nil {return _faed ;};_bgb :=_aefg .Ury ;for _ ,_caea :=range _gga {_baec :=_caea ._fdcce ;_dgf :=_gaa ._egcf .NewRectangle (_baec .Llx ,_bgb -_baec .Lly ,_baec .Urx -_baec .Llx ,-(_baec .Ury -_baec .Lly ));_dgf .SetFillColor (_bga );_dgf .SetBorderWidth (_fbca );
_dgf .SetFillOpacity (_egbf );if _fda :=_gaa ._egcf .Draw (_dgf );_fda !=nil {return nil ;};};};_gaa ._egcf .SetOutlineTree (_gaa ._egf .GetOutlineTree ());return nil ;};func _gage (_bbe _fe .PdfObject ,_bbb *_c .PdfFont )(string ,error ){_feec ,_addb :=_fe .GetStringBytes (_bbe );
if !_addb {return "",_fe .ErrTypeError ;};_dea :=_bbb .BytesToCharcodes (_feec );_gdd ,_efe ,_egb :=_bbb .CharcodesToStrings (_dea );if _egb > 0{_ee .Log .Debug ("\u0072\u0065nd\u0065\u0072\u0054e\u0078\u0074\u003a\u0020num\u0043ha\u0072\u0073\u003d\u0025\u0064\u0020\u006eum\u004d\u0069\u0073\u0073\u0065\u0073\u003d%\u0064",_efe ,_egb );
};_gab :=_gc .Join (_gdd ,"");return _gab ,nil ;};func (_cebb *regexMatcher )match (_agea string )([]*matchedIndex ,error ){_dfa :=_cebb ._beag .Pattern ;if _dfa ==nil {return nil ,_a .New ("\u006e\u006f\u0020\u0070at\u0074\u0065\u0072\u006e\u0020\u0063\u006f\u006d\u0070\u0069\u006c\u0065\u0064");
};var (_gbae =_dfa .FindAllStringIndex (_agea ,-1);_cfbf []*matchedIndex ;);for _ ,_ggag :=range _gbae {_cfbf =append (_cfbf ,&matchedIndex {_cec :_ggag [0],_cagg :_ggag [1],_fdg :_agea [_ggag [0]:_ggag [1]]});};return _cfbf ,nil ;};

// RedactRectanglePropsNew return a new pointer to a default RectangleProps object.
func RedactRectanglePropsNew ()*RectangleProps {return &RectangleProps {FillColor :_gd .ColorBlack ,BorderWidth :0.0,FillOpacity :1.0};};

// RedactionTerm holds the regexp pattern and the replacement string for the redaction process.
type RedactionTerm struct{Pattern *_ga .Regexp ;};func (_fdec *Redactor )redactPage (_edaa *_d .ContentStreamOperations ,_egce *_c .PdfPageResources )([]matchedBBox ,*_d .ContentStreamOperations ,error ){_cfd ,_fge :=_f .NewFromContents (_edaa .String (),_egce );
if _fge !=nil {return nil ,nil ,_fge ;};_edae ,_ ,_ ,_fge :=_cfd .ExtractPageText ();_edaa =_edae .GetContentStreamOps ();if _fge !=nil {return nil ,nil ,_fge ;};_ggad :=_edae .Marks ();_aeca :=_edae .Text ();_aed :=[]matchedBBox {};_fgbf :=make (map[_fe .PdfObject ][]localSpanMarks );
for _ ,_ggdc :=range _fdec ._cfc .Terms {_cbca ,_bfb :=_dag (_ggdc );if _bfb !=nil {return nil ,nil ,_bfb ;};_afg ,_bfb :=_cbca .match (_aeca );if _bfb !=nil {return nil ,nil ,_bfb ;};_afgg :=_beeac (_afg );for _ecb ,_bgbe :=range _afgg {_acf :=[]matchedBBox {};
for _ ,_bcba :=range _bgbe {_abee ,_dee ,_cce :=_egd (_bcba ,_ggad ,_ecb );if _cce !=nil {return nil ,nil ,_cce ;};_ffdf :=_ecdcd (_abee );for _gbaa ,_ceb :=range _ffdf {_feb :=localSpanMarks {_gbde :_ceb ,_afdd :_gbaa ,_cfbc :_ecb };_baf ,_ :=_dce (_ceb );
if _fdcb ,_eag :=_fgbf [_baf ];_eag {_fgbf [_baf ]=append (_fdcb ,_feb );}else {_fgbf [_baf ]=[]localSpanMarks {_feb };};};_acf =append (_acf ,_dee );};_aed =append (_aed ,_acf ...);};};_fge =_fd (_edaa ,_fgbf );if _fge !=nil {return nil ,nil ,_fge ;};
return _aed ,_edaa ,nil ;};

// Redactor represtents a Redactor object.
type Redactor struct{_egf *_c .PdfReader ;_cfc *RedactionOptions ;_egcf *_gd .Creator ;_cbcf *RectangleProps ;};

// New instantiates a Redactor object with given PdfReader and `regex` pattern.
func New (reader *_c .PdfReader ,opts *RedactionOptions ,rectProps *RectangleProps )*Redactor {if rectProps ==nil {rectProps =RedactRectanglePropsNew ();};return &Redactor {_egf :reader ,_cfc :opts ,_egcf :_gd .New (),_cbcf :rectProps };};

// RectangleProps defines properties of the redaction rectangle to be drawn.
type RectangleProps struct{FillColor _gd .Color ;BorderWidth float64 ;FillOpacity float64 ;};func _eae (_cf *_f .TextMarkArray )int {_dg :=0;_dd :=_cf .Elements ();if _dd [0].Text =="\u0020"{_dg ++;};if _dd [_cf .Len ()-1].Text =="\u0020"{_dg ++;};return _dg ;
};func _egd (_cfbcg []int ,_gdb *_f .TextMarkArray ,_dbb string )(*_f .TextMarkArray ,matchedBBox ,error ){_egca :=matchedBBox {};_bea :=_cfbcg [0];_cbf :=_cfbcg [1];_febe :=len (_dbb )-len (_gc .TrimLeft (_dbb ,"\u0020"));_dcgg :=len (_dbb )-len (_gc .TrimRight (_dbb ,"\u0020\u000a"));
_bea =_bea +_febe ;_cbf =_cbf -_dcgg ;_dbb =_gc .Trim (_dbb ,"\u0020\u000a");_gddg ,_ggc :=_gdb .RangeOffset (_bea ,_cbf );if _ggc !=nil {return nil ,_egca ,_ggc ;};_agb ,_bed :=_gddg .BBox ();if !_bed {return nil ,_egca ,_e .Errorf ("\u0073\u0070\u0061\u006e\u004d\u0061\u0072\u006bs\u002e\u0042\u0042ox\u0020\u0068\u0061\u0073\u0020\u006eo\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062\u006f\u0078\u002e\u0020s\u0070\u0061\u006e\u004d\u0061\u0072\u006b\u0073=\u0025\u0073",_gddg );
};_egca =matchedBBox {_fefff :_dbb ,_fdcce :_agb };return _gddg ,_egca ,nil ;};func _beeac (_dbbb []*matchedIndex )map[string ][][]int {_afad :=make (map[string ][][]int );for _ ,_cgc :=range _dbbb {_addg :=_cgc ._fdg ;_ddc :=[]int {_cgc ._cec ,_cgc ._cagg };
if _bcg ,_ebbd :=_afad [_addg ];_ebbd {_afad [_addg ]=append (_bcg ,_ddc );}else {_afad [_addg ]=[][]int {_ddc };};};return _afad ;};func _deg (_cdeg []localSpanMarks )(map[string ][]localSpanMarks ,[]string ){_ggf :=make (map[string ][]localSpanMarks );
_bae :=[]string {};for _ ,_dceg :=range _cdeg {_gca :=_dceg ._cfbc ;if _gce ,_da :=_ggf [_gca ];_da {_ggf [_gca ]=append (_gce ,_dceg );}else {_ggf [_gca ]=[]localSpanMarks {_dceg };_bae =append (_bae ,_gca );};};return _ggf ,_bae ;};func _fd (_de *_d .ContentStreamOperations ,_fb map[_fe .PdfObject ][]localSpanMarks )error {for _aa ,_ae :=range _fb {if _aa ==nil {continue ;
};_ea ,_bd ,_gb :=_ecdd (_de ,_aa );if !_gb {_ee .Log .Debug ("Pd\u0066\u004fb\u006a\u0065\u0063\u0074\u0020\u0025\u0073\u006e\u006ft\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0073\u0069\u0064\u0065\u0020\u0074\u0068\u0065\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0073\u0074r\u0065a\u006d\u0020\u006f\u0070\u0065\u0072\u0061\u0074i\u006fn\u0020\u0025s",_aa ,_de );
return nil ;};if _ea .Operand =="\u0054\u006a"{_ce :=_fgb (_ea ,_aa ,_ae );if _ce !=nil {return _ce ;};}else if _ea .Operand =="\u0054\u004a"{_gcb :=_dbe (_ea ,_aa ,_ae );if _gcb !=nil {return _gcb ;};}else if _ea .Operand =="\u0027"||_ea .Operand =="\u0022"{_dec :=_bf (_de ,_ea .Operand ,_bd );
if _dec !=nil {return _dec ;};_dec =_fgb (_ea ,_aa ,_ae );if _dec !=nil {return _dec ;};};};return nil ;};func _fbg (_gcfc *_f .TextMarkArray )string {_bcc :="";for _ ,_ebe :=range _gcfc .Elements (){_bcc +=_ebe .Text ;};return _bcc ;};func _ecdcd (_cdeb *_f .TextMarkArray )[]*_f .TextMarkArray {_ffb :=_cdeb .Elements ();
_eddc :=len (_ffb );var _ffc _fe .PdfObject ;_cba :=[]*_f .TextMarkArray {};_gbg :=&_f .TextMarkArray {};_gffg :=-1;for _bgde ,_ebf :=range _ffb {_faa :=_ebf .DirectObject ;_gffg =_ebf .Index ;if _faa ==nil {_fba :=_gdcg (_cdeb ,_bgde ,_gffg );if _ffc !=nil {if _fba ==-1||_fba > _bgde {_cba =append (_cba ,_gbg );
_gbg =&_f .TextMarkArray {};};};}else if _faa !=nil &&_ffc ==nil {if _gffg ==0&&_bgde > 0{_cba =append (_cba ,_gbg );_gbg =&_f .TextMarkArray {};};}else if _faa !=nil &&_ffc !=nil {if _faa !=_ffc {_cba =append (_cba ,_gbg );_gbg =&_f .TextMarkArray {};
};};_ffc =_faa ;_gbg .Append (_ebf );if _bgde ==(_eddc -1){_cba =append (_cba ,_gbg );};};return _cba ;};func _ca (_abe string ,_ggg *_c .PdfFont )[]byte {_ecg ,_fgf :=_ggg .StringToCharcodeBytes (_abe );if _fgf !=0{_ee .Log .Debug ("\u0057\u0041\u0052\u004e\u003a\u0020\u0073\u006fm\u0065\u0020\u0072un\u0065\u0073\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0065\u006e\u0063\u006f\u0064\u0065d\u002e\u000a\u0009\u0025\u0073\u0020\u002d\u003e \u0025\u0076",_abe ,_ecg );
};return _ecg ;};type matchedBBox struct{_fdcce _c .PdfRectangle ;_fefff string ;};

// Write writes the content of `re.creator` to writer of type io.Writer interface.
func (_ecgf *Redactor )Write (writer _ag .Writer )error {return _ecgf ._egcf .Write (writer )};type placeHolders struct{_db []int ;_cc string ;_dc float64 ;};func _dbe (_dgg *_d .ContentStreamOperation ,_ceg _fe .PdfObject ,_gff []localSpanMarks )error {_bdd ,_cdf :=_fe .GetArray (_dgg .Params [0]);
_eg :=[]_fe .PdfObject {};if !_cdf {_ee .Log .Debug ("\u0045\u0052\u0052OR\u003a\u0020\u0054\u004a\u0020\u006f\u0070\u003d\u0025s\u0020G\u0065t\u0041r\u0072\u0061\u0079\u0056\u0061\u006c\u0020\u0066\u0061\u0069\u006c\u0065\u0064",_dgg );return _e .Errorf ("\u0073\u0070\u0061\u006e\u004d\u0061\u0072\u006bs\u002e\u0042\u0042ox\u0020\u0068\u0061\u0073\u0020\u006eo\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062\u006f\u0078\u002e\u0020s\u0070\u0061\u006e\u004d\u0061\u0072\u006b\u0073=\u0025\u0073",_dgg );
};_bb ,_cef :=_deg (_gff );if len (_cef )==1{_fef :=_cef [0];_fc :=_bb [_fef ];if len (_fc )==1{_age :=_fc [0];_aaa :=_age ._gbde ;_eed :=_gdcb (_aaa );_deca ,_cdg :=_gage (_ceg ,_eed );if _cdg !=nil {return _cdg ;};_fa ,_cdg :=_aec (_age ,_aaa ,_eed ,_deca ,_fef );
if _cdg !=nil {return _cdg ;};for _ ,_dgc :=range _bdd .Elements (){if _dgc ==_ceg {_eg =append (_eg ,_fa ...);}else {_eg =append (_eg ,_dgc );};};}else {_fbe :=_fc [0]._gbde ;_ac :=_gdcb (_fbe );_ba ,_dbeg :=_gage (_ceg ,_ac );if _dbeg !=nil {return _dbeg ;
};_faf ,_dbeg :=_fdcc (_ba ,_fc );if _dbeg !=nil {return _dbeg ;};_edf :=_edbga (_faf );_ge :=_ccc (_ba ,_edf ,_ac );for _ ,_egc :=range _bdd .Elements (){if _egc ==_ceg {_eg =append (_eg ,_ge ...);}else {_eg =append (_eg ,_egc );};};};_dgg .Params [0]=_fe .MakeArray (_eg ...);
}else if len (_cef )> 1{_fg :=_gff [0];_ff :=_fg ._gbde ;_ ,_ddd :=_dce (_ff );_edd :=_ff .Elements ()[_ddd ];_cb :=_edd .Font ;_edb ,_baa :=_gage (_ceg ,_cb );if _baa !=nil {return _baa ;};_gfc ,_baa :=_fdcc (_edb ,_gff );if _baa !=nil {return _baa ;};
_add :=_edbga (_gfc );_aaf :=_ccc (_edb ,_add ,_cb );for _ ,_ffg :=range _bdd .Elements (){if _ffg ==_ceg {_eg =append (_eg ,_aaf ...);}else {_eg =append (_eg ,_ffg );};};_dgg .Params [0]=_fe .MakeArray (_eg ...);};return nil ;};func _fgb (_cdef *_d .ContentStreamOperation ,_adc _fe .PdfObject ,_cdc []localSpanMarks )error {var _ecd *_fe .PdfObjectArray ;
_dab ,_daa :=_deg (_cdc );if len (_daa )==1{_fafa :=_daa [0];_cag :=_dab [_fafa ];if len (_cag )==1{_ada :=_cag [0];_dca :=_ada ._gbde ;_agg :=_gdcb (_dca );_fgd ,_bg :=_gage (_adc ,_agg );if _bg !=nil {return _bg ;};_cfb ,_bg :=_aec (_ada ,_dca ,_agg ,_fgd ,_fafa );
if _bg !=nil {return _bg ;};_ecd =_fe .MakeArray (_cfb ...);}else {_fed :=_cag [0]._gbde ;_gcc :=_gdcb (_fed );_dcc ,_gdca :=_gage (_adc ,_gcc );if _gdca !=nil {return _gdca ;};_dddc ,_gdca :=_fdcc (_dcc ,_cag );if _gdca !=nil {return _gdca ;};_cdgb :=_edbga (_dddc );
_ace :=_ccc (_dcc ,_cdgb ,_gcc );_ecd =_fe .MakeArray (_ace ...);};}else if len (_daa )> 1{_fce :=_cdc [0];_acg :=_fce ._gbde ;_ ,_gbf :=_dce (_acg );_bgg :=_acg .Elements ()[_gbf ];_bba :=_bgg .Font ;_dggc ,_fea :=_gage (_adc ,_bba );if _fea !=nil {return _fea ;
};_aca ,_fea :=_fdcc (_dggc ,_cdc );if _fea !=nil {return _fea ;};_fafg :=_edbga (_aca );_dda :=_ccc (_dggc ,_fafg ,_bba );_ecd =_fe .MakeArray (_dda ...);};_cdef .Params [0]=_ecd ;_cdef .Operand ="\u0054\u004a";return nil ;};func _dce (_cff *_f .TextMarkArray )(_fe .PdfObject ,int ){var _bc _fe .PdfObject ;
_dge :=-1;for _be ,_fbd :=range _cff .Elements (){_bc =_fbd .DirectObject ;_dge =_be ;if _bc !=nil {break ;};};return _bc ,_dge ;};func _gdcg (_dgga *_f .TextMarkArray ,_gfe int ,_eagf int )int {_gcac :=_dgga .Elements ();_aeb :=_gfe -1;_cdb :=_gfe +1;
_fcg :=-1;if _aeb >=0{_fbed :=_gcac [_aeb ];_beea :=_fbed .ObjString ;_bge :=len (_beea );_dfg :=_fbed .Index ;if _dfg +1< _bge {_fcg =_aeb ;return _fcg ;};};if _cdb < len (_gcac ){_bfg :=_gcac [_cdb ];_bef :=_bfg .ObjString ;if _bef [0]!=_bfg .Text {_fcg =_cdb ;
return _fcg ;};};return _fcg ;};

// WriteToFile writes the redacted document to file specified by `outputPath`.
func (_ecf *Redactor )WriteToFile (outputPath string )error {if _ecddg :=_ecf ._egcf .WriteToFile (outputPath );_ecddg !=nil {return _e .Errorf ("\u0066\u0061\u0069l\u0065\u0064\u0020\u0074o\u0020\u0077\u0072\u0069\u0074\u0065\u0020t\u0068\u0065\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0066\u0069\u006c\u0065");
};return nil ;};type regexMatcher struct{_beag RedactionTerm };func _ebbg (_ebdd ,_ebee string )[]int {if len (_ebee )==0{return nil ;};var _fccf []int ;for _gbc :=0;_gbc < len (_ebdd );{_eda :=_gc .Index (_ebdd [_gbc :],_ebee );if _eda < 0{return _fccf ;
};_fccf =append (_fccf ,_gbc +_eda );_gbc +=_eda +len (_ebee );};return _fccf ;};func _dag (_bbd RedactionTerm )(*regexMatcher ,error ){return &regexMatcher {_beag :_bbd },nil };func _bfe (_bbf *_f .TextMarkArray )(float64 ,error ){_cge ,_bgda :=_bbf .BBox ();
if !_bgda {return 0.0,_e .Errorf ("\u0073\u0070\u0061\u006e\u004d\u0061\u0072\u006bs\u002e\u0042\u0042ox\u0020\u0068\u0061\u0073\u0020\u006eo\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062\u006f\u0078\u002e\u0020s\u0070\u0061\u006e\u004d\u0061\u0072\u006b\u0073=\u0025\u0073",_bbf );
};_gde :=_eae (_bbf );_bgdaf :=0.0;_ ,_abf :=_dce (_bbf );_abed :=_bbf .Elements ()[_abf ];_gcgg :=_abed .Font ;if _gde > 0{_bgdaf =_cbd (_gcgg ,_abed );};_cdca :=(_cge .Urx -_cge .Llx );_cdca =_cdca +_bgdaf *float64 (_gde );Tj :=_ecdc (_cdca ,_abed .FontSize ,_abed .Th );
return Tj ,nil ;};type replacement struct{_ec string ;_gag float64 ;_ed int ;};