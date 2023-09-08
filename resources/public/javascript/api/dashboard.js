/*
      dashboard.js will implement the endpoint required for
      the /api/v1/admin/dashboard? which returns information
      about information we should show

      /mainframe.js required
*/

// dashboard will make said api call then return the object found within it
async function dashboard() {
      const object = currentTokenInsideURL()

      /* saveReference is where we shall absorb the output into and handle accordingly */
      var saveReference = []

      // makes the fetch request towards the api endpoint
      await fetch(object + "/api/pages/dashboard").then(res => res.json()).then(res => {
            saveReference = res
      })

      return saveReference
}