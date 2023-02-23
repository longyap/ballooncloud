<template>
   

<!-- Sidebar -->
<section>
<ul class="pb-5 justify-left px-10 py-5 grid">
<l1 class="">
    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
  <path stroke-linecap="round" stroke-linejoin="round" d="M5.25 14.25h13.5m-13.5 0a3 3 0 01-3-3m3 3a3 3 0 100 6h13.5a3 3 0 100-6m-16.5-3a3 3 0 013-3h13.5a3 3 0 013 3m-19.5 0a4.5 4.5 0 01.9-2.7L5.737 5.1a3.375 3.375 0 012.7-1.35h7.126c1.062 0 2.062.5 2.7 1.35l2.587 3.45a4.5 4.5 0 01.9 2.7m0 0a3 3 0 01-3 3m0 3h.008v.008h-.008v-.008zm0-6h.008v.008h-.008v-.008zm-3 6h.008v.008h-.008v-.008zm0-6h.008v.008h-.008v-.008z" />
</svg >
<h3 class="font-bold text-2xl">Compute</h3>

<nav class="py-10 flex flex-col md:grid-cols-2">
    <a href="#" class="hover:text-blue-600 py-2">Overview</a>
    <a href="#" class="hover:text-blue-600 py-2">Instance</a>
    <a href="#" class="hover:text-blue-600 py-2">Custom Images</a>

</nav>
</l1 >
</ul>
<!-- Page Content -->
<div class="py-5" style="margin-left:10%">
   
 
    <h1 class="font-bold text-3xl"> Instances</h1>
    <span class="text-gray-600">
An instance is a compute host. Choose between virtual machines (VMs) and bare metal instances. The image that you use to launch an instance determines its operating system and other software.
</span>
 
<l1 class="grid grid-cols-4 py-5">

    <button type="button" onclick="location.href='/compute/create'"  class="py-2.5 px-5 mr-2 text-sm font-medium text-white bg-gradient-to-r from-blue-500 via-blue-600 to-blue-700 hover:bg-gradient-to-br inline-flex items-center">
    Create Instance
</button>
    <button type="button" class="py-2.5 px-5 mr-2 text-sm font-medium text-gray-900 bg-white border border-gray-200 hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-4 focus:outline-none focus:ring-blue-700 focus:text-blue-700 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700 inline-flex items-center">
    Table Setting
</button>

</l1>

    <div class=" shadow-md ">
    <table id="vmtablelist" class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
        <thead class="text-xs text-gray-700 uppercase bg-gray-50">
            <tr>
                <th scope="col" class="px-8 py-3">
                     Name
                </th>
                <th scope="col" class="px-8 py-3">
                    state
                </th>
                <th scope="col" class="px-8 py-3">
                    VCPU count  
                </th>
                <th scope="col" class="px-8 py-3">
                     Memory (GB)
                </th>
                <th scope="col" class="px-8 py-3">
                    UUID
                </th>
                <th scope="col" class="px-8 py-3">
                    Created Date
                </th>
                <th scope="col" class="px-8 py-3">
                    Private IP
                    </th>
                <th scope="col" class="px-8 py-3">
                    Public IP
                </th>
                <th scope="col" class="px-8 py-3">
                    Shape
                </th>
            </tr>
        </thead>
        <tbody id="tbody">
 
        </tbody>
    </table>
</div>

</div>
  
</section>
</template>

<script>
fetch('http://localhost:8080/api/vm')
  .then(response => response.json())
  .then(data => {
   // const thead = document.querySelector('thead');
    const tbody = document.getElementById('tbody');
    for (const vm of data) {
      const row = document.createElement('tr');
        row.classList.add("border-b");
      const nameCell = document.createElement('th');
      nameCell.textContent = vm.name;
      const stateCell = document.createElement('th');
      stateCell.textContent = vm.state;
      const vcpusCell = document.createElement('th');
      vcpusCell.textContent = vm.vcpus;
      const ramCell = document.createElement('th');
      ramCell.textContent = vm.ram;
      const uuidCell = document.createElement('th');
      uuidCell.textContent = vm.uuid;
      const dateCell = document.createElement('th');
      dateCell.textContent = vm.date;
      row.appendChild(nameCell);
      row.appendChild(stateCell);
      row.appendChild(vcpusCell);
      row.appendChild(ramCell);
      row.appendChild(uuidCell);
      row.appendChild(dateCell);
      tbody.appendChild(row);
    }
  })
  
  .catch(error => console.error(error));
</script>