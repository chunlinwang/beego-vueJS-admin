<template>
  <menu-detail :is-edit="true"/>
</template>

<script>
import MenuDetail from './components/MenuDetail'

export default {
  name: 'EditForm',
  components: { MenuDetail }
}
</script>

