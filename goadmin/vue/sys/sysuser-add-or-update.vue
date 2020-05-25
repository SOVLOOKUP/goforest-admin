<template>
    <el-dialog
            :title="!dataForm.id ? '新增' : '修改'"
            :close-on-click-modal="false"
            :visible.sync="visible">
        <el-form :model="dataForm" :rules="dataRule" ref="dataForm" label-width="80px">
            <el-form-item label="用户名" prop="Username">
                <el-input v-model="dataForm.Username" placeholder="用户名"></el-input>
            </el-form-item>
            <el-form-item label="密码" prop="Password">
                <el-input v-model="dataForm.Password" placeholder="密码"></el-input>
            </el-form-item>
            <el-form-item label="秘钥" prop="Salt">
                <el-input v-model="dataForm.Salt" placeholder="秘钥"></el-input>
            </el-form-item>
            <el-form-item label="手机号" prop="Mobile">
                <el-input v-model="dataForm.Mobile" placeholder="手机号"></el-input>
            </el-form-item>
            <el-form-item label="创建时间" prop="CreateTime">
                <el-input v-model="dataForm.CreateTime" placeholder="创建时间"></el-input>
            </el-form-item>
            <el-form-item label="更新时间" prop="UpdateTime">
                <el-input v-model="dataForm.UpdateTime" placeholder="更新时间"></el-input>
            </el-form-item>
            <el-form-item label="状态  0：禁用   1：正常" prop="Status">
                <el-input v-model="dataForm.Status" placeholder="状态  0：禁用   1：正常"></el-input>
            </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
      <el-button @click="visible = false">取消</el-button>
      <el-button type="primary" @click="dataFormSubmit()">确定</el-button>
    </span>
    </el-dialog>
</template>

<script>
    export default {
        data() {
            return {
                visible: false,
                dataForm: {
            Id
        :
            0,
            Username
        :
            '',
            Password
        :
            '',
            Salt
        :
            '',
            Mobile
        :
            '',
            CreateTime
        :
            '',
            UpdateTime
        :
            '',
            Status
        :
            '',
        },
            dataRule: {
                Username
            :
                [
                    {required: true, message: '用户名不能为空', trigger: 'blur'}
                ],
                Password
            :
                [
                    {required: true, message: '密码不能为空', trigger: 'blur'}
                ],
                Salt
            :
                [
                    {required: true, message: '秘钥不能为空', trigger: 'blur'}
                ],
                Mobile
            :
                [
                    {required: true, message: '手机号不能为空', trigger: 'blur'}
                ],
                CreateTime
            :
                [
                    {required: true, message: '创建时间不能为空', trigger: 'blur'}
                ],
                UpdateTime
            :
                [
                    {required: true, message: '更新时间不能为空', trigger: 'blur'}
                ],
                Status
            :
                [
                    {required: true, message: '状态  0：禁用   1：正常不能为空', trigger: 'blur'}
                ],
            }
        }
        },
        methods: {
            init(id) {
                this.dataForm.$
                {.
                    pk.AttrName
                }
                = id || 0
                this.visible = true
                this.$nextTick(() => {
                    this.$refs['dataForm'].resetFields()
                    if (this.dataForm.Id
                )
                    {
                        this.$http({
                            url: this.$http.adornUrl(`/sys/sysuser/get/${this.dataForm.Id}`),
                            method: 'get',
                            params: this.$http.adornParams()
                        }).then(({data}) => {
                            if (data && data.code === 0) {
                                this.dataForm.$
                                {
                                    $column.AttrName
                                }
                                = data.$
                                {
                                    $.classname
                                }
                            .
                                $
                                {
                                    $column.AttrName
                                }
                                this.dataForm.$
                                {
                                    $column.AttrName
                                }
                                = data.$
                                {
                                    $.classname
                                }
                            .
                                $
                                {
                                    $column.AttrName
                                }
                                this.dataForm.$
                                {
                                    $column.AttrName
                                }
                                = data.$
                                {
                                    $.classname
                                }
                            .
                                $
                                {
                                    $column.AttrName
                                }
                                this.dataForm.$
                                {
                                    $column.AttrName
                                }
                                = data.$
                                {
                                    $.classname
                                }
                            .
                                $
                                {
                                    $column.AttrName
                                }
                                this.dataForm.$
                                {
                                    $column.AttrName
                                }
                                = data.$
                                {
                                    $.classname
                                }
                            .
                                $
                                {
                                    $column.AttrName
                                }
                                this.dataForm.$
                                {
                                    $column.AttrName
                                }
                                = data.$
                                {
                                    $.classname
                                }
                            .
                                $
                                {
                                    $column.AttrName
                                }
                                this.dataForm.$
                                {
                                    $column.AttrName
                                }
                                = data.$
                                {
                                    $.classname
                                }
                            .
                                $
                                {
                                    $column.AttrName
                                }
                            }
                        })
                    }
                })
            },
            // 表单提交
            dataFormSubmit() {
                this.$refs['dataForm'].validate((valid) => {
                    if (valid) {
                        this.$http({
                            url: this.$http.adornUrl(`/sys/sysuser/${!this.dataForm.Id ? 'save' : 'update'})`),
                            method: 'post',
                            data: this.$http.adornData({
                                'Id': this.dataForm.$
                        {
                            $column.AttrName
                        }
                    ||
                        undefined,
                            'Username'
                    :
                        this.dataForm.$
                        {
                            $column.AttrName
                        }
                    ,
                            'Password'
                    :
                        this.dataForm.$
                        {
                            $column.AttrName
                        }
                    ,
                            'Salt'
                    :
                        this.dataForm.$
                        {
                            $column.AttrName
                        }
                    ,
                            'Mobile'
                    :
                        this.dataForm.$
                        {
                            $column.AttrName
                        }
                    ,
                            'CreateTime'
                    :
                        this.dataForm.$
                        {
                            $column.AttrName
                        }
                    ,
                            'UpdateTime'
                    :
                        this.dataForm.$
                        {
                            $column.AttrName
                        }
                    ,
                            'Status'
                    :
                        this.dataForm.$
                        {
                            $column.AttrName
                        }
                    ,
                    })
                    }).
                        then(({data}) => {
                            if (data && data.code === 0) {
                                this.$message({
                                    message: '操作成功',
                                    type: 'success',
                                    duration: 1500,
                                    onClose: () => {
                                        this.visible = false
                                        this.$emit('refreshDataList')
                                    }
                                })
                            } else {
                                this.$message.error(data.msg)
                            }
                        })
                    }
                })
            }
        }
    }
</script>