<div class="row">
    <div class="col-md-3">
        <div class="tree">
            <ul style="padding-left: 10px;">
                {{{.docsTree}}}
            </ul>
        </div>
    </div>
    <div id="ArticleCentent" class="col-md-9" style="border: lightgray 1px solid;background-color: #f5f5f5;border-radius: 3px;height: 500px">
        {{{.centent}}}
    </div>
</div>
<script src="/tree/index.js"></script>
<div id="editDocs" class="row">
    <!-- 模态框（Modal） -->
<div class="modal fade" id="myModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
	<div class="modal-dialog">
		<div class="modal-content">
			<div class="modal-header">
				<button type="button" class="close" data-dismiss="modal" aria-hidden="true">
					&times;
				</button>
				<h4 class="modal-title" id="myModalLabel">
					开发文档树操作
				</h4>
			</div>
			<div class="modal-body">
                <h2 id="docName"></h2>
                <label for="name">添加子文档树名字:</label><input name="docname" type="text"/><button>添加</button>
                <br/>
                <button id="dtldoc">删除当前文档树</button>
                <br/>
                <button id="updoc">上移当前文档树</button>
                <br/>
                <button id="downdox">下移当前文档树</button>
			</div>
			<div class="modal-footer">
				<button type="button" class="btn btn-default" data-dismiss="modal">关闭
				</button>
			</div>
		</div><!-- /.modal-content -->
	</div><!-- /.modal -->
</div>
    <script src="/js/editDocs.js"></script>
</div>