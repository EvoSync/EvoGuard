let date = new Date();
let hour = date.getHours();

// WelcomeSelector will decide via the time on what greetings to show
function WelcomeSelector(id) {
    if (hour >= 5 && hour < 12) {
        msg = "Good morning, "; 
    } else if (hour >= 12 && hour < 17) {
        msg = "Good afternoon, "
    } else if (hour >= 17 && hour < 23) {
        msg = "Good evening, "
    } else if (hour >= 23 && hour < 5) {
        msg = "Good night, "
    } else {
        msg = "Welcome,"
    }

    document.getElementById(id).innerHTML = msg;
}

WelcomeSelector("timeSpecifier");