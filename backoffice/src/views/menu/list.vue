<template>
  <div class="app-container">

    <el-table v-loading="listLoading" :data="list" border fit highlight-current-row style="width: 100%">
      <el-table-column :label="$t('menu.id')" align="center" >
        <template slot-scope="scope">
          <span>{{ scope.row._id }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('menu.number')" align="center" >
        <template slot-scope="scope">
          <span>{{ scope.row.name }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('menu.price')" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.price }} &euro;</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('menu.on_sale')" align="center" >
        <template slot-scope="scope">
          <span>{{ scope.row.on_sale }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('menu.price_on_sale')" align="center" >
        <template slot-scope="scope">
          <span>{{ scope.row.price_on_sale }} &euro;</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('menu.actions')" align="center">
        <template slot-scope="scope">
          <router-link :to="'/menu/edit/'+scope.row._id">
            <el-button type="primary" size="small" icon="el-icon-edit">{{$t('edit')}}</el-button>
          </router-link>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit" @pagination="getList" />

  </div>
</template>

<script>
import { fetchList } from '@/api/menu'
import Pagination from '@/components/Pagination' // Secondary package based on el-pagination

export default {
  name: 'MenuList',
  components: { Pagination },
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'info',
        deleted: 'danger'
      }
      return statusMap[status]
    }
  },
  data() {
    return {
      list: null,
      total: 0,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 20
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      fetchList(this.listQuery).then(response => {
        this.list = response.data.items
        this.total = response.data.total
        this.listLoading = false
      })
    },
    handleSizeChange(val) {
      this.listQuery.limit = val
      this.getList()
    },
    handleCurrentChange(val) {
      this.listQuery.page = val
      this.getList()
    }
  }
}
</script>

<style scoped>
.edit-input {
  padding-right: 100px;
}
.cancel-btn {
  position: absolute;
  right: 15px;
  top: 10px;
}
</style>
