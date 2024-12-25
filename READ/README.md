# Hangman-Web
 
Le projet présenté est un site web nous permettant de visualiser le jeu du pendu mais aussi d'y jouer. Le jeu consiste à deviner un mot en un nombre limité de tentatives en proposant des lettres ou en devinant directement le mot.
Prérequis
## To run
`commande a mettre pour lancer le code`

```
PS C:\Users\ghass\hangman-web> go run webmotif.go words.txt
```
 
Pour exécuter ce programme, vous aurez besoin d'un environnement Go installé sur votre machine.
Ensuite seul le lien du site web sera nécessaire pour pouvoir jouer.
 
Fonctionnalités
 
Le jeu du Pendu est un jeu de devinette de mots.
Voici comment le programme fonctionne :
 
    Ici vous vous trouver sur la page web du jeu.
    Si vous cliquez sur jouer il se passera ceci :
    Le programme charge un fichier contenant une liste de mots (fichier words.txt) et en sélectionne un au hasard.
    Le joueur a un nombre limité de tentatives pour deviner le mot ou proposer des lettres.
    Si le joueur devine correctement le mot ou toutes les lettres du mot, il gagne.
    Sinon, le jeu se termine lorsque le nombre de tentatives atteint zéro.
 
Le programme gère également la répétition de lettres et affiche une représentation visuelle du pendu en fonction du nombre de tentatives restantes.
Commandes du jeu
 
Le joueur peut interagir avec le jeu de la manière suivante :
 
    Le joueur peut proposer une lettre en entrant une seule lettre majuscule.
    Si le joueur entre le mot entier (en majuscules), le jeu vérifie si le mot est correct.
    Si le joueur entre une lettre déjà proposée, le jeu l'en informe et demande de proposer une autre lettre.
 
Auteur :
Ce programme a été écrit par Amine, Dany, Kayss, Matis
 
Note : Pour jouer au jeu du Pendu, assurez vous que le fichier words.txt contenant une liste de mots est présent dans le même répertoire que le programme. Vous pouvez ajouter vos propres mots à ce fichier pour personnaliser le jeu.