<template>
  <div id="search">
    <div class="container">
      <div class="row">
        <div class="col-7">
             <a-input-search placeholder="input search text" style="width: 200px" @search="onSearch" />
          <!-- <div class="input-group mb-3">
            <input
              type="text"
              class="form-control rounded-pill border-0"
              placeholder="Recipient's username"
              aria-label="🔍"
              aria-describedby="button-addon2"
              style="rounded-pill"
            />
           
               <a-button type="primary" shape="circle">
      A
    </a-button> -->
            
          <!-- </div> -->
        </div>
      </div>
    </div>
  </div>
</template>


<style>
#search {
  min-height: 667px;
  background-color: #edeeed;
}
</style>