package language

import "fmt"

var MessageCommandStart = []translate{
	translate{ Lang: "en", Text: "Welcome! Subscribe trackers to receive notifications when the registries are opened." },
	translate{ Lang: "es", Text: "¡Bienvenido! Suscríbete a trackers para recibir una notificación cuando los registros estén abiertos." },
}

var MessageCommandUnknown = []translate{
	translate{ Lang: "en", Text: "Command unknown. Try with /info." },
	translate{ Lang: "es", Text: "Comando desconocido. Prueba con /info para obtener ayuda." },
}

var MessageCommandInfo = []translate{
	translate{ Lang: "en", Text: fmt.Sprintf("This bot monitor private torrent trackers to notify you when the registries are opened. Can obtain further information on %s.", RepoUrl) },
	translate{ Lang: "es", Text: fmt.Sprintf("Este bot monitoriza trackers torrent privados para avisarte cuando abran los registros. Puedes obtener más información en %s.", RepoUrl) },
}