<template>
  <page-detail :is-edit="true"/>
</template>

<script>
import PageDetail from './components/PageDetail'

export default {
  name: 'EditForm',
  components: { PageDetail }
}
</script>

