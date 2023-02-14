<script setup>
</script>

<template>
  <main>
    <div class="grid grid-cols-6 gap-4">
      <div class="col-start-2 col-span-3">
        <h1 class="px-6 py-2 text-3xl font-bold">
    Dashboard
  </h1>
      </div>
      <button class="px-2 py-2 text-sm bg-red-800 hover:bg-slate-200 text-slate-50" @click="increment">test</button>


    </div>
    <hr>
   
    <div class="flex p-4 mb-4 text-sm text-blue-800 rounded-lg bg-blue-50 dark:bg-gray-800 dark:text-blue-400" role="alert">
  <svg aria-hidden="true" class="flex-shrink-0 inline w-5 h-5 mr-3" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd"></path></svg>
  <span class="sr-only">Info</span>
  <div>
    <span class="font-medium">Ensure that these requirements are met:</span>
      <ul class="mt-1.5 ml-4  list-inside">
        <li>At least 10 characters (and up to 100 characters)</li>

    </ul>
  </div>
</div>
  

 </main>
</template>
