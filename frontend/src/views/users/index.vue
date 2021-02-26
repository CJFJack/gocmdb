<template>
    <d2-container>
        <template slot="header">用户管理</template>
        <template>
            <el-table
                    :data="tableData"
                    style="width: 100%">
                <el-table-column v-for="col in tableColumns"
                                 :prop="col.key"
                                 :label="col.title">
                </el-table-column>
            </el-table>
        </template>
    </d2-container>
</template>

<script>
    import api from "../../api";

    export default {
        name: "users",
        data() {
            return {
                tableData: [],
                tableColumns: [],
            }
        },
        methods: {
            async listTable() {
                const res = await this.$api.LIST_USERS()
                this.tableData = res.tableData
                this.tableColumns = res.tableColumns
            }
        },
        mounted() {
            this.listTable()
        }
    }
</script>

<style scoped>

</style>