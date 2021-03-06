package language

import "fmt"

var PhraseChooseSubscription = []translate{
	translate{ Lang: "en", Text: "What do you want to subscribe to?" },
	translate{ Lang: "es", Text: "¿A que deseas suscribirte?" },
}

var PhraseIDontUnderstand = []translate{
	translate{ Lang: "en", Text: "Sorry, I couldn't understand you" },
	translate{ Lang: "es", Text: "Lo siento, pero no te he entendido" },
}

var PhraseInsertLanguage = []translate{
	translate{ Lang: "en", Text: "What languages do you like subscribe to? (Send the ISO code, example: en for English. Set before the symbol - to unsubscribe, example: -en). The available languages are:" },
	translate{ Lang: "es", Text: "¿A que idioma te quieres suscribir? (Envía el ISO code, ejemplo: es para Español. Pon un - delante para borrar la suscripción, ejemplo: -es). Los idiomas disponibles son: " },
}

var PhraseLanguageSubscribed = []translate{
	translate{ Lang: "en", Text: "Great! You are now subscribed to the trackers for this language. You can refine the notifications by subscribing to more specific categories." },
	translate{ Lang: "es", Text: "¡Bien! Te has suscrito a los trackers de este idioma. Puedes ajustar más las notificaciones suscribiendote a categorías concretas." },
}

var PhraseInsertCategory = []translate{
	translate{ Lang: "en", Text: "What categories do you like subscribe to? (Send the name, example: movies. Set before the symbol - to unsubscribe, example: -movies). The available categories are:" },
	translate{ Lang: "es", Text: "¿A que categoría te quieres suscribir? (Envía el nombre, ejemplo: movies. Pon un - delante para borrar la suscripción, ejemplo: -movies). Las categorías disponibles son: " },
}

var PhraseCategorySubscribed = []translate{
	translate{ Lang: "en", Text: "Great! You are now subscribed to the trackers of this category. You can refine the notifications by subscribing to more specific categories." },
	translate{ Lang: "es", Text: "¡Bien! Ahora estás suscrito a los trackers de esta categoría. Puedes afinar más las notificacions suscribiendote a un lenguaje o lenguajes concretos." },
}

var PhraseInternalError = []translate{
	translate{ Lang: "en", Text: fmt.Sprintf("Wow, something went wrong. Tell us what you were doing in %s to see if we can help you.", IssueUrl) },
	translate{ Lang: "es", Text: fmt.Sprintf("Wow, algo ha ido mal. Cuentanos que estábas haciendo en %s a ver si podemos ayudarte.", IssueUrl) },
}

var PhraseNoSubscriptions = []translate{
	translate{ Lang: "en", Text: "Looks like you haven't subscribed to any notification yet. Try the /subscribe command." },
	translate{ Lang: "es", Text: "Parece que aún no te has suscrito a ninguna notificación. Prueba con el comando /subscribe." },
}

var PhraseUserSubscriptionsCategories = []translate{
	translate{ Lang: "en", Text: "You are currently subscribed to the categories: %s." },
	translate{ Lang: "es", Text: "Estás suscrito a las siguientes categorías: %s." },
}

var PhraseUserSubscriptionsLanguages = []translate{
	translate{ Lang: "en", Text: "You are currently subscribed to the languages: %s." },
	translate{ Lang: "es", Text: "Estás suscrito a las siguientes idiomas: %s." },
}

var PhraseCategoryUnrecognized = []translate{
	translate{ Lang: "en", Text: "I'm sorry, I can't recognize this category." },
	translate{ Lang: "es", Text: "Lo siento, no reconozco esa categoría." },
}

var PhraseLanguageUnrecognized = []translate{
	translate{ Lang: "en", Text: "I'm sorry, I can't recognize this language." },
	translate{ Lang: "es", Text: "Lo siento, no reconozco este idioma." },
}

var PhraseLanguageUnsubscribed = []translate{
	translate{ Lang: "en", Text: "Ok, I'll stop warning you about trackers of this language." },
	translate{ Lang: "es", Text: "Vale, dejaré de avisarte sobre trackers de este idioma." },
}

var PhraseCategoryUnsubscribed = []translate{
	translate{ Lang: "en", Text: "Ok, I'll stop warning you about trackers of this category." },
	translate{ Lang: "es", Text: "Vale, dejaré de avisarte sobre trackers de esta categoría." },
}

var PhraseNoSubscribed = []translate{
	translate{ Lang: "en", Text: "Ups, looks like we didn't already have any of your subscriptions." },
	translate{ Lang: "es", Text: "Ups, parece que ya no teníamos ninguna suscripción tuya." },
}

var PhraseActionsCanceled = []translate{
	translate{ Lang: "en", Text: "All actions cancelled." },
	translate{ Lang: "es", Text: "Se han cancelado todas las acciones." },
}

var PhraseTrackerOnline = []translate{
	translate{ Lang: "en", Text: "Hey! It seems that the %s tracker registrations are open. Try to signing up: %s." },
	translate{ Lang: "es", Text: "¡Ey! Parece que los registros del tracker %s estan abiertos. Intenta registrarte: %s." },
}

var PhraseTrackerMaybeOnline = []translate{
	translate{ Lang: "en", Text: "Hi, I've detected that the %s tracker could have opened the registrations. Try to signing up: %s." },
	translate{ Lang: "es", Text: "Hola, he detectado que el tracker %s podría haber abierto los registros. Intenta registrarte: %s." },
}
