<h1> ja goddag

Her er lige en super hurtigt guide til platformio det er nemlig reeet nemt at arbejde med.

Den måde det hele fungerer på er at vi har 2 seperate platformio projekter uart og motor.

UART projektet er bindeledet mellem raspberry pi og arduinoen som benytter uart + x10.

Motor projektet er arduinoen som har til ansvar at modtage og sende signaler via x10 og så styre motoren som driver låsen.

For at platformio ikke bliver vred så skal man sørge for at åbne projektet i den rigtige mappe, det vil sige at hvis man gerne vil arbejde med uart arduinoen så skal man altså åbne projektet fra uart mappen af, og det samme er gældende med motor.

Det betyder at man skal gå ind i: semesterprojekt2->arduino->uart

Hvis man er en terminal kriger så hedder det
```cd ./arduino/uart && code .```
med udgangspunkt at man allerede er inde i semesterprojekt2 mappen.