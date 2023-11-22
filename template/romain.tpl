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
        window.location.href = `/romain?${urlParams.toString()}`;
      } else if (formationParam === "cyber") {
        urlParams.set("formation", "dev");
        // redirect to the new url
        window.location.href = `/romain?${urlParams.toString()}`;
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
        const url = `/romain?${urlParams.toString()}`;

        // redirect to the new url
        window.location.href = url;
      }

</script>
<div class="container mx-auto mt-10">
  <div class="wrapper bg-white rounded shadow w-full">
    <div class="header flex justify-between border-b p-2">
      <div>
        <span class="text-lg font-bold">Calendrier de Romain pour {{ .Month.MonthName }} {{ .Year }}  </span>
        <button id="previous" class="p-1" onclick="changeMonth('previous')">
          <svg width="1em" fill="gray" height="1em" viewBox="0 0 16 16" class="bi bi-arrow-left-circle" fill="currentColor" xmlns="http://www.w3.org/2000/svg">
            <path fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
            <path fill-rule="evenodd" d="M8.354 11.354a.5.5 0 0 0 0-.708L5.707 8l2.647-2.646a.5.5 0 1 0-.708-.708l-3 3a.5.5 0 0 0 0 .708l3 3a.5.5 0 0 0 .708 0z" />
            <path fill-rule="evenodd" d="M11.5 8a.5.5 0 0 0-.5-.5H6a.5.5 0 0 0 0 1h5a.5.5 0 0 0 .5-.5z" />
          </svg>
        </button>
        <button id="next" class="p-1" onclick="changeMonth('next')">
          <svg width="1em" fill="gray" height="1em" viewBox="0 0 16 16" class="bi bi-arrow-right-circle" fill="currentColor" xmlns="http://www.w3.org/2000/svg">
            <path fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
            <path fill-rule="evenodd" d="M7.646 11.354a.5.5 0 0 1 0-.708L10.293 8 7.646 5.354a.5.5 0 1 1 .708-.708l3 3a.5.5 0 0 1 0 .708l-3 3a.5.5 0 0 1-.708 0z" />
            <path fill-rule="evenodd" d="M4.5 8a.5.5 0 0 1 .5-.5h5a.5.5 0 0 1 0 1H5a.5.5 0 0 1-.5-.5z" />
          </svg>
        </button>
      </div>
    </div>

    {{$CurrentDay := .CurrentDay}}


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
            {{ if .Empty }}
            <!-- If the day is empty, display nothing in the header -->
            &nbsp;
            {{ else }}
            {{ .DayName }}
            {{ end }}
          </th>
          {{ end }}
        </tr>
      </thead>
      <tbody>
        <tr>
          {{ range $day := .Days }}  
          {{ if and (eq $day.DayDate.Year $CurrentDay.Year) (eq $day.DayDate.Month $CurrentDay.Month) (eq $day.DayDate.Day $CurrentDay.Day) }}
            <!-- Code when the date part of $day.DayDate is equal to the date part of $CurrentDay -->
            <td
                class="p-1 border cursor-pointer duration-500 ease h-40 lg:w-30 md:w-30 bg-gray-300 overflow-auto sm:w-20 transition w-10 xl:w-40"
            >
                {{ range $event := $day.DayEvents }}
                    <div
                        class="{{ if $day.CDG57 }}bg-green-400{{ else }}bg-purple-400{{ end }} text-white rounded p-1 text-sm mb-1"
                    >
                        <span class="event-name">{{ $event.EventName }}</span>
                        <br>
                        <span class="time">{{ $event.EventStart.Format "15:04" }}</span> 
                        <span class="time">{{ $event.EventEnd.Format "15:04" }}</span>
                    </div>
                {{ end }}
            </td>
          {{ else }}
            <td
            class="p-1 border cursor-pointer duration-500 ease h-40 lg:w-30 md:w-30 overflow-auto sm:w-20 transition w-10 xl:w-40"
            >
                {{ range $event := $day.DayEvents }}
                    <div
                        class="{{ if $day.CDG57 }}bg-green-400{{ else }}bg-purple-400{{ end }} text-white rounded p-1 text-sm mb-1"
                    >
                        <span class="event-name">{{ $event.EventName }}</span>
                        <br>
                        <span class="time">{{ $event.EventStart.Format "15:04" }}</span> 
                        <span class="time">{{ $event.EventEnd.Format "15:04" }}</span>
                    </div>
                {{ end }}
            </td>
          {{ end }}
        {{ end }}
           
        </tr>
      </tbody>
    </table>
    {{ end }}
    {{ end }}
  </div>
</div>
