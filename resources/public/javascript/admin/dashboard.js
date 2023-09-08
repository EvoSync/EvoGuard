/* 
      dashboard.js will implement the api routing required, this
      includes how we absorb the incoming api routine and render
      onto the homepage
*/

// once the action is complete, we continue to update all values we
// want to display in/on the dashboard.
dashboard().then(res => {
      document.getElementById("Username").textContent = res["username"]
      document.getElementById("TotalMembers").textContent = res["totalMembers"]
      document.getElementById("TotalActivatedLicenses").textContent = res["totalActivatedLicenses"]
      document.getElementById("TotalPurchasesToday").textContent = res["totalPurchasesToday"]
      document.getElementById("NumberApplications").textContent = res["numberApplications"]

      // Disables the spinner and proceeds to show the dashboard
      document.getElementById("spinner").style.visibility = "hidden"
      document.getElementById("app").style.visibility = "visible"

      // Creates the charts we wish to display on the dashboard
      new Chart("WeeklySales", {type: "bar", data: {labels: res["weeklySalesFields"], datasets: [{backgroundColor: "#00A3FF", data: res["weeklySalesValues"]}]}, options: {legend: {display: false}}});
      new Chart("ApplicationSlices", {type: "pie",data: {labels: res["applications"],datasets: [{backgroundColor: res["applicationsColours"],data: res["applicationsValues"]}]},options: {legend: {display: false}, title: {display: false}}});

})