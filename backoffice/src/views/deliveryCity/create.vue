<template>
  <delivery-city-detail :is-edit="false"/>
</template>

<script>
import DeliveryCityDetail from './components/DeliveryCityDetail'

export default {
  name: 'CreateForm',
  components: { DeliveryCityDetail }
}
</script>

