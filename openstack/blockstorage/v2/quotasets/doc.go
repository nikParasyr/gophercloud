/*
Package quotasets enables retrieving and managing Block Storage quotas.

Example to Get a Quota Set

	quotaset, err := quotasets.Get(context.TODO(), blockStorageClient, "project-id").Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", quotaset)

Example to Get Quota Set Usage

	quotaset, err := quotasets.GetUsage(context.TODO(), blockStorageClient, "project-id").Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", quotaset)

Example to Update a Quota Set

	updateOpts := quotasets.UpdateOpts{
		Volumes: gophercloud.IntToPointer(100),
	}

	quotaset, err := quotasets.Update(context.TODO(), blockStorageClient, "project-id", updateOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", quotaset)

Example to Update a Quota set with volume_type quotas

	updateOpts := quotasets.UpdateOpts{
		Volumes: gophercloud.IntToPointer(100),
		Extra: map[string]any{
			"gigabytes_foo": gophercloud.IntToPointer(100),
			"snapshots_foo": gophercloud.IntToPointer(10),
			"volumes_foo":   gophercloud.IntToPointer(10),
		},
	}

	quotaset, err := quotasets.Update(context.TODO(), blockStorageClient, "project-id", updateOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", quotaset)

Example to Delete a Quota Set

	err := quotasets.Delete(context.TODO(), blockStorageClient, "project-id").ExtractErr()
	if err != nil {
		panic(err)
	}
*/
package quotasets
