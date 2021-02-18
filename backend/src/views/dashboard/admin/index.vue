<template>
  <div class="dashboard-editor-container">

    <!--<panel-group />-->

    <el-row :gutter="5" class="background-white">
      <el-col :xs="{span: 24}" :sm="{span: 24}" :md="{span: 24}" :lg="{span: 14}" :xl="{span: 12}" style="padding-right:8px;margin-bottom:30px;">
        <AgencyConfigForm />
      </el-col>
      <el-col :xs="{span: 24}" :sm="{span: 24}" :md="{span: 24}" :lg="{span: 10}" :xl="{span: 12}" style="padding-right:8px;margin-bottom:30px;">
        <transaction-table />
      </el-col>
    </el-row>
  </div>
</template>

<script>
// import PanelGroup from './components/PanelGroup'
import TransactionTable from './components/TransactionTable'
import AgencyConfigForm from '../admin/components/AgencyConfigForm'

export default {
  name: 'DashboardAdmin',
  components: {
    // PanelGroup,
    TransactionTable,
    AgencyConfigForm
  },
  data() {
    return {}
  },
  methods: {}
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
.dashboard-editor-container {
  padding: 32px;
  background-color: rgb(240, 242, 245);
  .chart-wrapper {
    background: #fff;
    padding: 16px 16px 0;
    margin-bottom: 32px;
  }
}
</style>
