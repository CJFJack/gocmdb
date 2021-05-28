<template>
    <d2-container>
        <template slot="header">Prometheus - Job管理</template>
        <el-button type="primary" @click="add('add')">新增</el-button>
        <template>
            <el-table
                    :data="tableData"
                    v-loading="loading"
                    size="mini"
                    style="width: 100%">
                <el-table-column width="1">
                </el-table-column>
                <div v-for="col in tableColumns">
                    <el-table-column :prop="col.key" :label="col.title" v-if="col.key==='Node'">
                        <template slot-scope="scope">
                            <div v-html="scope.row[col.key].Hostname"></div>
                        </template>
                    </el-table-column>
                    <el-table-column :prop="col.key" :label="col.title" v-else-if="col.key==='CreatedAt'">
                        <template slot-scope="scope">
                            <div v-html="formatDateTime(scope.row[col.key])"></div>
                        </template>
                    </el-table-column>
                    <el-table-column :prop="col.key" :label="col.title" v-else-if="col.key==='UpdatedAt'">
                        <template slot-scope="scope">
                            <div v-html="formatDateTime(scope.row[col.key])"></div>
                        </template>
                    </el-table-column>
                    <el-table-column :prop="col.key" :label="col.title" v-else>
                    </el-table-column>
                </div>
                <el-table-column
                        fixed="right"
                        label="操作"
                        width="180">
                    <template slot-scope="scope">
                        <el-button @click="modify(scope.row)" type="success" size="mini">修改</el-button>
                        <el-button type="danger" @click="delConfirm(scope.row)" size="mini">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <br>
            <div class="block">
                <el-pagination
                        @size-change="paginationSizeChange"
                        @current-change="paginationCurrentChange"
                        :current-page="pagination.currentPage"
                        :page-sizes="[10, 20, 50, 100, 200]"
                        :page-size="pagination.pageSize"
                        layout="total, sizes, prev, pager, next, jumper"
                        :total="pagination.total"
                        :prev-text="pagination.prevText"
                        :next-text="pagination.nextText">
                </el-pagination>
            </div>
        </template>

        <el-dialog title="新增 / 编辑" :visible.sync="dialogVisible" :destroy-on-close="true">
            <el-form ref="form" :model="form" :rules="rules" label-width="120px">
                <el-form-item label="Node" prop="Node">
                    <el-select v-model="form.Node">
                        <el-option v-for="item in nodeSelectOptions" :value="item.ID" :key="item.ID" :label="item.Hostname"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="Key" prop="Key">
                    <el-input v-model="form.Key" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="备注" prop="Remark">
                    <el-input type="text" v-model="form.Remark" autocomplete="off"></el-input>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="dialogVisible = false">取 消</el-button>
                <el-button type="primary" @click="saveForm">确 定</el-button>
            </div>
        </el-dialog>
    </d2-container>
</template>

<script>
    export default {
        name: "prometheus_job",
        data() {
            return {
                tableData: [],
                tableColumns: [],
                tableTotal: 0,
                // 分页设置
                pagination: {
                    currentPage: 1,
                    pageSize: 10,
                    total: 100,
                    prevText: '上一页',
                    nextText: '下一页',
                    pageSizes: [5, 10, 20, 30, 40, 50, 100],
                    layout: 'sizes, prev, pager, next, jumper, ->, total, slot'
                },
                dialogVisible: false,
                form: {
                    Node: '',
                    Key: '',
                    Remark: '',
                },
                rules: {
                    Node: [{required: true, message: '请选择节点', trigger: 'blur'}],
                    Key: [{required: true, message: '请输入Key', trigger: 'blur'}],
                    Remark: [{required: true, message: '请输入备注', trigger: 'blur'}],
                },
                action: '',
                loading: false,
                nodeSelectOptions: [],
            }
        },
        methods: {
            notice(title, message) {
                this.$alert(message, title, {
                    confirmButtonText: '确定',
                });
            },

            formatDateTime(value) {
                if (value) {
                    return value.replace("T", " ").replace("+08:00", "")
                }
                return ""
            },

            async listTable() {
                this.loading = true
                const data = {
                    pagination: {
                        currentPage: this.pagination.currentPage,
                        pageSize: this.pagination.pageSize,
                    }
                }
                const res = await this.$api.LIST_PROMETHEUS_JOBS(data)
                this.tableData = res.tableData
                this.tableColumns = res.tableColumns
                this.pagination.total = res.tableTotal
                this.nodeSelectOptions = res.nodeSelectOptions
                this.loading = false
            },

            saveForm() {
                this.$refs["form"].validate(async (valid) => {
                    if (valid) {
                        if (this.action === 'add') {
                            await this.$api.ADD_PROMETHEUS_JOB(this.form)
                            this.listTable()
                        } else if (this.action === 'modify') {
                            const res = await this.$api.MODIFY_PROMETHEUS_JOB(this.form)
                            // this.notice("恭喜", "修改成功")
                            this.listTable()
                        }
                        this.dialogVisible = false
                    } else {
                        console.log('error submit!!');
                        return false;
                    }
                });
            },

            add(action) {
                this.form = {
                    Node: '',
                    Key: '',
                    Remark: '',
                }
                this.action = action
                this.dialogVisible = true
            },

            modify(row) {
                this.action = 'modify'
                this.form = JSON.parse(JSON.stringify(row))
                this.form.Node = this.form.Node.ID
                this.dialogVisible = true
            },

            async del(row) {
                let data = {
                    id: row.ID
                }
                await this.$api.DEL_PROMETHEUS_JOB(data)
                this.$message({
                    type: 'success',
                    message: '删除成功!'
                });
                this.listTable()
            },

            delConfirm(row) {
                this.$confirm('此操作将删除Job, 是否继续?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    this.del(row)
                }).catch(() => {
                    this.$message({
                        type: 'info',
                        message: '已取消删除'
                    });
                });
            },

            // 当分页改变时触发
            paginationCurrentChange(currentPage) {
                this.pagination.currentPage = currentPage
                this.listTable()
            },

            // 分页选择器改变后触发的函数
            paginationSizeChange(pageSize) {
                this.pagination.pageSize = pageSize
                this.listTable()
            },

        },
        mounted() {
            this.listTable()
        }
    }
</script>

<style scoped>

</style>
