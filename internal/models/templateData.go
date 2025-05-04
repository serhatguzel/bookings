package models

type TemplateData struct {
	StringMap map[string]string      // Eğer sadece string göndermek istiyorsanız bu şekilde yapabilirsiniz.
	IntMap    map[string]int         // Eğer sadece int göndermek istiyorsanız bu şekilde yapabilirsiniz.
	FloatMap  map[string]float32     // Eğer sadece float göndermek istiyorsanız bu şekilde yapabilirsiniz.
	Data      map[string]interface{} // Eğer birden fazla veri göndermek istiyorsanız bu şekilde yapabilirsiniz.
	CSRFToken string                 // CSRF token göndermek istiyorsanız bu şekilde yapabilirsiniz.
	Flash     string                 // Flash mesajı göndermek istiyorsanız bu şekilde yapabilirsiniz.
	Warning   string                 // Uyarı mesajı göndermek istiyorsanız bu şekilde yapabilirsiniz.
	Error     string                 // Hata mesajı göndermek istiyorsanız bu şekilde yapabilirsiniz.
}
