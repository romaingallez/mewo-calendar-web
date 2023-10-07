<script>
  function changeFormation() {
    // get formation from url
    const urlParams = new URLSearchParams(window.location.search);
    const formationParam = urlParams.get("formation");
    console.log(formationParam);
    // if formationParam is not empty
    if (formationParam != null) {

      // if formationParam is dev then change to cyber and vice versa
      if (formationParam === "dev") {
        urlParams.set("formation", "cyber");
        // redirect to the new url
        window.location.href = `/month?${urlParams.toString()}`;
      } else if (formationParam === "cyber") {
        urlParams.set("formation", "dev");
        // redirect to the new url
        window.location.href = `/month?${urlParams.toString()}`;
      }
   
    } 
    // create the new url
    

  }
  function changeMonth(direction) {
        // get current date
        const date = new Date();

        // get month param from url
        const urlParams = new URLSearchParams(window.location.search);
        const monthParam = urlParams.get("month");
        const yearParam = urlParams.get("year");

        // if month param is not empty, use that instead of the current month
        let month;
        if (monthParam != null) {
          month = parseInt(monthParam, 10);
        } else {
          month = date.getMonth() + 1;
        }
        let year;
        if (yearParam != null) {
          year = parseInt(yearParam, 10);
        } else {
          year = date.getFullYear();
        }

        // if direction is next, increment the month, otherwise decrement it
        if (direction === "next") {
          month++;
        } else if (direction === "previous") {
          month--;
        }

        // if month is greater than 12, set it to 1 and increment the year
        if (month > 12) {
          month = 1;
          year++;
        }

        // if month is less than 1, set it to 12 and decrement the year
        if (month < 1) {
          month = 12;
          year--;
        }

        // add a 0 if the month is less than 10
        if (month < 10) {
          month = "0" + month;
        }

        // update the month and year url params
        urlParams.set("month", month);
        urlParams.set("year", year);

        // create the new url
        const url = `/month?${urlParams.toString()}`;

        // redirect to the new url
        window.location.href = url;
      }

</script>
<div class="container mx-auto mt-10">
  <div class="wrapper bg-white rounded shadow w-full">
    <div class="header justify-between p-2">
      <span class="font-bold text-lg"></span>
      <button
        id="change"
        class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
        onclick="changeFormation()"
      >
        Change to {{ .InvertFormation }}
      </button>

      <div class="buttons float-right">
        <button id="previous" class="p-1" onclick="changeMonth('previous')">
          <!-- SVG content for left arrow -->
        </button>
        <button id="next" class="p-1" onclick="changeMonth('next')">
          <!-- SVG content for right arrow -->
        </button>
      </div>

      <div class="relative max-w-sm">
        <div
          class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none"
        ></div>
      </div>
    </div>

    {{ range .Month.Weeks }}
    {{ $hasEvents := false }}
    {{ range .Days }}
    {{ if .DayEvents }}
    {{ $hasEvents = true }}
    {{ end }}
    {{ end }}

    {{ if $hasEvents }}
    <table class="w-full">
      <thead>
        <tr>
          {{ range .Days }}
          <th
            class="text-gray-600 font-normal text-sm py-2 px-2 border-b border-gray-200"
          >
            {{ .DayName }} {{ .DayDate.Day }} {{ .DayDate.Month }} {{ .DayDate.Year }}
          </th> 
          {{ end }}
        </tr>
      </thead>
      <tbody>
        <tr>
          {{ range $day := .Days }}
          <td
            class="p-1 border cursor-pointer duration-500 ease h-40 hover:bg-gray-300 lg:w-30 md:w-30 overflow-auto sm:w-20 transition w-10 xl:w-40"
          >
            {{ range $event := $day.DayEvents }}
            <div
              class="event bg-purple-400 text-white rounded p-1 text-sm mb-1"
            >
              <span class="event-name">{{ $event.EventName }}</span>
              <br>
              <span class="time">{{  $event.EventStart.Format "15:04" }}</span> <span class="time"> {{ $event.EventEnd.Format "15:04" }}</span>
            </div>
            {{ end }}
          </td>
          {{ end }}
        </tr>
      </tbody>
    </table>
    {{ end }}
    {{ end }}
  </div>
</div>
z