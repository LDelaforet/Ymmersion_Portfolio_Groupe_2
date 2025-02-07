// au chargement de la page
document.addEventListener('DOMContentLoaded', function() {
    // Recuperation des objets du HTML
    const cards = document.querySelectorAll('.card');
    const audio = document.getElementById('son');
    const overlay = document.querySelector('.overlay');
    const closeButton = document.querySelector('.close-button');

    let activeCard = null;

    // Toggle la classe clicked sur la carte envoyée
    function retournerCarte(idCarte) {
        const carte = document.getElementById(idCarte);
        if (carte) {
            carte.classList.toggle('clicked');
        }
    }

    // Retourne les cartes basées sur l'id de la carte cliquée
    function GetFlipCard(currentCardID) {
        switch (currentCardID) {
            case "1":
                retournerCarte("1");
                break;
            case "2":
                retournerCarte("2");
                retournerCarte("5");
                break;
            case "3":
                retournerCarte("3");
                retournerCarte("1");
                break;
            case "4":
                retournerCarte("4");
                retournerCarte("2");
                break;
            case "5":
                retournerCarte("5");
                retournerCarte("6");
                break;
            case "6":
                retournerCarte("5");
                retournerCarte("6");
                break;
            default:
                break;
        }
    }

    cards.forEach(card => {
        // Quand une carte est cliquée 
        card.addEventListener('click', (event) => {
            const clickedCard = event.currentTarget;

            // Joue le son et retourne la carte
            if (audio) {
                audio.currentTime = 0;
                audio.play();
            }
            // Retourne la carte
            GetFlipCard(clickedCard.id);
            // Vérifie si toutes les cartes sont retournées
            checkAllCardsClicked();
        });
    });

    // Gestion du zoom des épreuves
    const epreuveCards = document.querySelectorAll('.epreuve-card');
    epreuveCards.forEach(card => {
        card.addEventListener('click', function() {
            zoomCard(card);
        });
    });

    function zoomCard(card) {
        console.log("Zoom sur la carte", card);
        activeCard = card;
        const zoomContent = card.querySelector('.zoom-content');
        if (!zoomContent) {
            console.error("Élément .zoom-content non trouvé dans", card);
            return;
        }
        card.classList.add('zoomed');
        overlay.classList.add('active');
        closeButton.classList.add('visible');
        setTimeout(() => {
            zoomContent.classList.add('visible');
        }, 500);
        document.body.style.overflow = 'hidden';
    }

    //Gestion de la fermeture du zoom
    function unzoomCard() {
        if (!activeCard) return;
        console.log("Unzoom de la carte", activeCard);
        const zoomContent = activeCard.querySelector('.zoom-content');
        if (zoomContent) {
            zoomContent.classList.remove('visible');
        }
        setTimeout(() => {
            activeCard.classList.remove('zoomed');
            overlay.classList.remove('active');
            closeButton.classList.remove('visible');
            activeCard = null;
            document.body.style.overflow = 'auto';
        }, 300);
    }

    //Gestion de la fermeture des carte d'epreuves
    closeButton.addEventListener('click', unzoomCard);
    overlay.addEventListener('click', unzoomCard);
    document.addEventListener('keydown', function(e) {
        if (e.key === 'Escape' && activeCard) {
            unzoomCard();
        }
    });
});


//Animation navabar lors du scroll
let header = document.querySelector("nav");
let menu = document.querySelectorAll("a");

document.addEventListener("scroll", () => {
    let scrollPosition = window.scrollY;
    if (scrollPosition > 0) {
        header.style.backgroundColor = "#CCC";
        header.style.color = "#000000";
        menu.forEach((item) => {
            item.style.color = "#000000";
        });
        header.style.transition = "background-color 1s ease-in-out, color 1s ease-in-out";
    } else {
        header.style.backgroundColor = "transparent";
        header.style.color = "#FFFFFF";
        header.style.borderBottom = "none";
        menu.forEach((item) => {
            item.style.color = "#FFFFFF";
        });
        header.style.transition = "background-color 1s ease-in-out, color 1s ease-in-out";
    }
});

/* NAVBAR DYNAMIQUE */
document.querySelectorAll('.nav-links a').forEach(anchor => {
    anchor.addEventListener('click', function (e) {
        e.preventDefault();

        // Récupère l'id de la section à atteindre
        const targetId = this.getAttribute('href').substring(1);
        const targetSection = document.getElementById(targetId);

        // Défilement fluide vers la section cible
        targetSection.scrollIntoView({ behavior: 'smooth' });
    });
});


