<template>
  <promo-code-detail :is-edit="true"/>
</template>

<script>
import PromoCodeDetail from './components/PromoCodeDetail'

export default {
  name: 'EditForm',
  components: { PromoCodeDetail }
}
</script>

