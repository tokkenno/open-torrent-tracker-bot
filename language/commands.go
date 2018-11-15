package language

import "fmt"

var MessageCommandStart = []translate{
	translate{ Lang: "en", Text: "Welcome! Subscribe to trackers in order to receive notifications once the registries are opened." },
	translate{ Lang: "es", Text: "¡Bienvenido! Suscríbete a trackers para recibir una notificación cuando los registros estén abiertos." },
}

var MessageCommandUnknown = []translate{
	translate{ Lang: "en", Text: "Unknown command . Try with /info." },
	translate{ Lang: "es", Text: "Comando desconocido. Prueba con /info para obtener ayuda." },
}

var MessageCommandInfo = []translate{
	translate{ Lang: "en", Text: fmt.Sprintf("This bot monitors private torrent trackers and will send you notifications once the registries are opened. Further information on %s.", RepoUrl) },
	translate{ Lang: "es", Text: fmt.Sprintf("Este bot monitoriza trackers torrent privados para avisarte cuando abran los registros. Puedes obtener más información en %s.", RepoUrl) },
}
