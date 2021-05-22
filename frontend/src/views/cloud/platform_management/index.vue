<template>
    <d2-container>
        <template slot="header">云平台管理</template>
        <el-button type="primary" @click="add('add')">新增</el-button>
        <template>
            <el-table
                    :data="tableData"
                    size="mini"
                    style="width: 100%">
                <el-table-column width="1">
                </el-table-column>
                <div v-for="col in tableColumns">
                    <el-table-column :prop="col.key" :label="col.title" v-if="col.key==='CreatedTime'">
                        <template slot-scope="scope">
                            <span>{{ formatDateTime(scope.row[col.key]) }}</span>
                        </template>
                    </el-table-column>
                    <el-table-column :prop="col.key" :label="col.title" v-else-if="col.key==='SyncTime'">
                        <template slot-scope="scope">
                            <span>{{ formatDateTime(scope.row[col.key]) }}</span>
                        </template>
                    </el-table-column>
                    <el-table-column :prop="col.key" :label="col.title" v-else-if="col.key==='Type'">
                        <template slot-scope="scope">
                            <span>{{ typeOptions[scope.row[col.key]] }}</span>
                        </template>
                    </el-table-column>
                    <el-table-column :prop="col.key" :label="col.title" v-else-if="col.key==='Status'">
                        <template slot-scope="scope">
                            <span>{{ statusOptions[scope.row[col.key]] }}</span>
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
                <el-form-item label="名称" prop="Name">
                    <el-input v-model="form.Name" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="类型" prop="Type">
                    <el-select v-model="form.Type">
                        <el-option v-for="(item, key) in typeOptions" :label="item" :key="key" :value="key"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="地址" prop="Addr">
                    <el-input v-model="form.Addr" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="AccessKey" prop="AccessKey">
                    <el-input type="password" v-model="form.AccessKey" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="SecretKey" prop="SecretKey">
                    <el-input type="password" v-model="form.SecretKey" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="地域" prop="Region">
                    <el-input v-model="form.Region" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="备注" prop="Remark">
                    <el-input type="textarea" v-model="form.Remark" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="状态">
                    <el-select v-model="form.Status" placeholder="请选择状态">
                        <el-option v-for="(text, key, index) in statusOptions" :label="text" :value="index"></el-option>
                    </el-select>
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
        name: "cloud_platform_management",
        watch: {
          "action": function (val) {
              if (val === 'add') {
                  this.rules.AccessKey = [{required: true, message: '请输入AccessKey', trigger: 'blur'}]
                  this.rules.SecretKey = [{required: true, message: '请输入SecretKey', trigger: 'blur'}]
              } else if (val === 'modify') {
                  this.rules.AccessKey = [{required: false}]
                  this.rules.SecretKey = [{required: false}]
              }
          }
        },
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
                    Name: '',
                    Type: '',
                    Addr: '',
                    AccessKey: "",
                    SecretKey: "",
                    Region: "",
                    Remark: "",
                    Status: 0,
                },
                rules: {
                    Name: [{required: true, message: '请输入名称', trigger: 'blur'}],
                    Type: [{required: true, message: '请输入类型', trigger: 'blur'}],
                    Addr: [{required: true, message: '请输入地址', trigger: 'blur'}],
                    AccessKey: [{required: true, message: '请输入AccessKey', trigger: 'blur'}],
                    SecretKey: [{required: true, message: '请输入SecretKey', trigger: 'blur'}],
                    Region: [{required: true, message: '请输入地域', trigger: 'blur'}],
                },
                typeOptions: {},
                statusOptions: {},
                action: '',
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
                const data = {
                    pagination: {
                        currentPage: this.pagination.currentPage,
                        pageSize: this.pagination.pageSize,
                    }
                }
                const res = await this.$api.LIST_CLOUD_PLATFORM(data)
                this.tableData = res.tableData
                this.tableColumns = res.tableColumns
                this.pagination.total = res.tableTotal
                this.typeOptions = res.typeOptions
                this.statusOptions = res.statusOptions
            },

            saveForm() {
                this.$refs["form"].validate(async (valid) => {
                    if (valid) {
                        if (this.action === 'add') {
                            await this.$api.ADD_CLOUD_PLATFORM(this.form)
                            this.listTable()
                        } else if (this.action === 'modify') {
                            const res = await this.$api.MODIFY_CLOUD_PLATFORM(this.form)
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
                    Name: '',
                    Type: '',
                    Addr: '',
                    AccessKey: "",
                    SecretKey: "",
                    Region: "",
                    Remark: "",
                    Status: 0,
                }
                this.action = action
                this.dialogVisible = true
            },

            modify(row) {
                this.action = 'modify'
                this.form = JSON.parse(JSON.stringify(row))
                this.dialogVisible = true
            },

            async del(row) {
                let data = {
                    id: row.ID
                }
                await this.$api.DEL_CLOUD_PLATFORM(data)
                this.$message({
                    type: 'success',
                    message: '删除成功!'
                });
                this.listTable()
            },

            delConfirm(row) {
                this.$confirm('此操作将删除云平台, 是否继续?', '提示', {
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
