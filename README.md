# Calendrier MEWO en ligne

Ce projet permet d'accéder à son calendrier MEWO en ligne, sans avoir à se connecter à un portail web ou à partager ses événements avec des personnes qui n'ont pas accès.

## Pourquoi ?

* Commodité : consulter son calendrier depuis n'importe quel appareil, sans avoir à ce connecter.
* Le challange technique

## Comment ?

Le fichier ICS du calendrier est converti en JSON et stocké dans une base de données MongoDB. Le site web accède au JSON via un proxy.

**Méthode :**

1. Le logiciel que le campus Mewo utilise ([SC-FORM](https://www.sc-form.com/)) permet d'exporter un fichier ICS du calendrier.
2. Ce fichier ICS est converti en JSON à l'aide d'un projet annexe.
3. Le JSON est stocké dans une base de données MongoDB.
4. Le site web accède au JSON via un proxy.

