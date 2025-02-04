function retournerCarte(idCarte) {
    const carte = document.getElementById(idCarte);
    if (carte) {
        carte.classList.toggle('clicked');
    }
}

// Finalement je le fais pas dans le go
function GetFlipCard(currentCardID) {
    switch (currentCardID) {
        case "4":
            retournerCarte("4");
            break;
        case "5":
            retournerCarte("5");
            retournerCarte("2");
            break;
        case "6":
            retournerCarte("6");
            retournerCarte("4");
            break;
        case "1":
            retournerCarte("1");
            retournerCarte("5");
            break;
        case "2":
            retournerCarte("2");
            retournerCarte("3");
            break;
        case "3":
            retournerCarte("2");
            retournerCarte("3");
            break;
        default:
            break;
    }
}

function checkAllCardsClicked() {
    const cards = document.querySelectorAll('.card');
    let allFlipped = true;

    cards.forEach(card => {
        if (!card.classList.contains('clicked')) {
            allFlipped = false;
        }
    });

    if (allFlipped) {
        alert("Toutes les cartes ont été cliquées!");
    }
}

document.querySelectorAll('.card').forEach(card => {
    card.addEventListener('click', (event) => {
        let sound = document.getElementById("son");
        const clickedCard = event.currentTarget;
        

        sound.currenTime = 0;
        sound.play();
        
        GetFlipCard(clickedCard.id);
        checkAllCardsClicked();
    });
});