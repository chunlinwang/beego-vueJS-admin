<template>
  <promo-code-detail :is-edit="false"/>
</template>

<script>
import PromoCodeDetail from './components/PromoCodeDetail'

export default {
  name: 'CreateForm',
  components: { PromoCodeDetail }
}
</script>

