<template>
    
<div class="bg-white border-gray-200 px-2 sm:px-4 py-2.5 rounded dark:bg-gray-900">

   

  <div class="container flex flex-wrap items-center justify-between mx-auto">
    <div class="grid grid-cols-2">
        <div>
      <button type="button" class="text-grey-800 hover:text-blue-600 font-bold text-sm px-5 py-2.5 text-center mr-2 mb-2">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
  <path stroke-linecap="round" stroke-linejoin="round" d="M9 15L3 9m0 0l6-6M3 9h12a6 6 0 010 12h-3" />
</svg>
    
        Back
    </button>
    </div>
   <div class="">
    <a href="#" class="flex items-center">
        <span class="self-center text-xl font-semibold whitespace-nowrap">Instance Name</span>
    </a>
    <ul class="list-disc pl-5">
    <li href="#" class="self-center">
        <span class=" text-green-700">Running</span>
    </li>
</ul>
   </div>
</div>
   
    <div class="hidden w-full md:block md:w-auto" id="navbar-default">
        <button type="button" class="text-grey-800 hover:text-blue-600 font-bold text-sm px-5 py-2.5 text-center mr-2 mb-2">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
      <path stroke-linecap="round" stroke-linejoin="round" d="M16.862 4.487l1.687-1.688a1.875 1.875 0 112.652 2.652L10.582 16.07a4.5 4.5 0 01-1.897 1.13L6 18l.8-2.685a4.5 4.5 0 011.13-1.897l8.932-8.931zm0 0L19.5 7.125M18 14v4.75A2.25 2.25 0 0115.75 21H5.25A2.25 2.25 0 013 18.75V8.25A2.25 2.25 0 015.25 6H10" />
      </svg>Edit
        </button>
        <button type="button" class="text-grey-800 hover:text-blue-600 font-bold rounded-lg text-sm px-5 py-2.5 text-center mr-2 mb-2">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
            <path stroke-linecap="round" stroke-linejoin="round" d="M5.25 5.653c0-.856.917-1.398 1.667-.986l11.54 6.348a1.125 1.125 0 010 1.971l-11.54 6.347a1.125 1.125 0 01-1.667-.985V5.653z" />
            </svg>
            Start
        </button>
        <button type="button" class="text-grey-800 hover:text-blue-600 font-bold rounded-lg text-sm px-5 py-2.5 text-center mr-2 mb-2">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
              <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 5.25v13.5m-7.5-13.5v13.5" />
            </svg>
            Pause   
        </button>
        <button type="button" class="text-grey-800 hover:text-blue-600 font-bold rounded-lg text-sm px-5 py-2.5 text-center mr-2 mb-2">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
            <path stroke-linecap="round" stroke-linejoin="round" d="M5.25 7.5A2.25 2.25 0 017.5 5.25h9a2.25 2.25 0 012.25 2.25v9a2.25 2.25 0 01-2.25 2.25h-9a2.25 2.25 0 01-2.25-2.25v-9z" />
            </svg>
            Stop
        </button>
        <button type="button" class="text-grey-800 hover:text-blue-600 font-bold rounded-lg text-sm px-5 py-2.5 text-center mr-2 mb-2">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
          <path stroke-linecap="round" stroke-linejoin="round" d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0l3.181 3.183a8.25 8.25 0 0013.803-3.7M4.031 9.865a8.25 8.25 0 0113.803-3.7l3.181 3.182m0-4.991v4.99" />
            </svg>
            Restart
        </button>

        <button type="button" class="text-grey-800 hover:text-blue-600 font-bold text-sm px-5 py-2.5 text-center mr-2 mb-2">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                <path stroke-linecap="round" stroke-linejoin="round" d="M6.827 6.175A2.31 2.31 0 015.186 7.23c-.38.054-.757.112-1.134.175C2.999 7.58 2.25 8.507 2.25 9.574V18a2.25 2.25 0 002.25 2.25h15A2.25 2.25 0 0021.75 18V9.574c0-1.067-.75-1.994-1.802-2.169a47.865 47.865 0 00-1.134-.175 2.31 2.31 0 01-1.64-1.055l-.822-1.316a2.192 2.192 0 00-1.736-1.039 48.774 48.774 0 00-5.232 0 2.192 2.192 0 00-1.736 1.039l-.821 1.316z" />
                <path stroke-linecap="round" stroke-linejoin="round" d="M16.5 12.75a4.5 4.5 0 11-9 0 4.5 4.5 0 019 0zM18.75 10.5h.008v.008h-.008V10.5z" />
            </svg>
            Create Snapshot
        </button>

        <button type="button" class="text-red-700 hover:text-white border border-red-700 hover:bg-red-800 font-bold text-sm px-5 py-2.5 text-center mr-2 mb-2">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
              <path stroke-linecap="round" stroke-linejoin="round" d="M14.74 9l-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 01-2.244 2.077H8.084a2.25 2.25 0 01-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 00-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 013.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 00-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 00-7.5 0" />
            </svg>
            Terminate
        </button>
        

    </div>
  </div>
</div>
<hr>

<section>
    
    <ul class="pb-5 justify-left px-10 py-5 grid">
    <h3 class="font-bold text-2xl">Resources</h3>    
<l1>
<nav class="py-10 flex flex-col md:grid-cols-2">
    <a href="#" class="hover:text-blue-600 py-2">Details</a>
    <a href="#" class="hover:text-blue-600 py-2">Metrics</a>
    <a href="#" class="hover:text-blue-600 py-2">Attached Block Volume</a>
    <a href="#" class="hover:text-blue-600 py-2">Attached VNIC</a>
    <a href="#" class="hover:text-blue-600 py-2">Boot Volume</a>
    <a href="#" class="hover:text-blue-600 py-2">Vnc Viewer</a>
    <a href="#" class="hover:text-blue-600 py-2">Console Logs</a>


</nav>
</l1 >
</ul>
<section>
<div class="p-4 sm:ml-12 sm:mr-12">
    <h1 class= "py-2 px-10 text-2xl text-slate-900">
    Details
  </h1>
</div>

</section>
</section>
</template>
<script>
//fetch('http://localhost:8080/api/vm/1?uuid='+{ uuid })
 // .then(response => response.json())
 // .then(data => {

 // })





</script>