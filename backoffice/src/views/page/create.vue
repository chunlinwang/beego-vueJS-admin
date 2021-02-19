<template>
  <page-detail :is-edit="false"/>
</template>

<script>
import PageDetail from './components/PageDetail'

export default {
  name: 'CreateForm',
  components: { PageDetail }
}
</script>

