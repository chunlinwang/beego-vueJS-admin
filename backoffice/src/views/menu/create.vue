<template>
  <menu-detail :is-edit="false"/>
</template>

<script>
import MenuDetail from './components/MenuDetail'

export default {
  name: 'CreateForm',
  components: { MenuDetail }
}
</script>

