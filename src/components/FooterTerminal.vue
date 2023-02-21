<template>
      <div id="footerterminal"  class="hidden bottom-navbar w-full">
    <div class="">
        <!-- Modal content -->
        <div class="p-4 bg-white shadow sm:p-5">
            <!-- Modal header -->
            <div class="flex justify-between items-center mb-4 rounded-t border-b sm:mb-5 dark:border-gray-600">

              <div class="text-sm font-medium text-center text-gray-500 border-b border-gray-200 dark:text-gray-400 dark:border-gray-700">
   
                <ul class="flex flex-wrap -mb-px">
                  <label for="tab1" class=" nav mr-2rounded-t-lg hover:text-gray-600 hover:bg-gray-200">
                  <div class="inline-block p-4 border-b-2 border-transparent ">Terminal
                  </div>
                </label>
                  <label for="tab2" class="nav mr-2 hover:text-gray-600 hover:bg-gray-200 rounded-t-lg">
                    <div class="inline-block p-4 border-b-2 border-transparent">Editor</div>  
                </label>
              
               </ul>
              </div>
              <button type="button" class="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm p-1.5 inline-flex items-center dark:hover:bg-gray-600 dark:hover:text-white" data-modal-toggle="updateProductModal">
                    <svg xmlns="http://www.w3.org/2000/svg" height="24" width="24" fill="currentColor"><path d="M5 21q-.825 0-1.413-.587Q3 19.825 3 19V5q0-.825.587-1.413Q4.175 3 5 3h7v2H5v14h14v-7h2v7q0 .825-.587 1.413Q19.825 21 19 21Zm4.7-5.3-1.4-1.4L17.6 5H14V3h7v7h-2V6.4Z"/></svg>  
                  <span class="sr-only">Close</span>
                  </button>
                <button type="button" @click="minimizeterminal" class="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm p-1.5 ml-auto inline-flex items-center dark:hover:bg-gray-600 dark:hover:text-white" data-modal-toggle="updateProductModal">
                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 12h-15" />
                  </svg>
                  </button>
                  <button type="button" @click="maximizeterminal" class="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm p-1.5 inline-flex items-center dark:hover:bg-gray-600 dark:hover:text-white" data-modal-toggle="updateProductModal">
                    <svg xmlns="http://www.w3.org/2000/svg"  height="24" width="24" fill="currentColor" viewBox="0 0 512 512"><title>Square</title><path d="M416 448H96a32.09 32.09 0 01-32-32V96a32.09 32.09 0 0132-32h320a32.09 32.09 0 0132 32v320a32.09 32.09 0 01-32 32z" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32"/></svg> <span class="sr-only">Close</span>
                  </button>
                  <button type="button" @click="closeterminal" class="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm p-1.5 inline-flex items-center dark:hover:bg-gray-600 dark:hover:text-white" data-modal-toggle="updateProductModal">
                    <svg aria-hidden="true" class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd"></path></svg>
                    <span class="sr-only">Close</span>
                  </button>
                </div>
            <!-- Modify pure css or with js later -->
                <div id=terminalform class="grid gap-4 mb-4 sm:grid-cols-2">
                  <input type="radio" class="hidden peer" name="tabs" id="tab1" checked />
                  <input type="radio" class="hidden peer" name="tabs" id="tab2" />
                  <div class="tab content1">Tab1 Contents</div>
                  <div class="tab content2">Tab2 Contents</div>
            
                  </div>

        </div>
    </div>
</div>
</template>
<script>
function minimizeterminal(){
  document.getElementById("terminalform").style.display="none";
}
function maximizeterminal(){
  document.getElementById("terminalform").style.display="block";
}
function closeterminal(){
  document.getElementById("footerterminal").style.display="none";
}

</script>