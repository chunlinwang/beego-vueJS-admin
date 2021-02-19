<template>
  <delivery-city-detail :is-edit="true"/>
</template>

<script>
import DeliveryCityDetail from './components/DeliveryCityDetail'

export default {
  name: 'EditForm',
  components: { DeliveryCityDetail }
}
</script>

