var ErrFoundCircle = fmt.Errorf("Found circle")

func dfs(rbac *RBAC, id string, skipped map[string]struct{}, stack []string) error {
	if _, ok := skipped[id]; ok {
		return nil
	}
	for _, item := range stack {
		if item == id {
			return ErrFoundCircle
		}
	}
	parents := rbac.parents[id]
	if len(parents) == 0 {
		stack = nil, skipped[id] = empty
		return nil
	}
	stack = append(stack, id)
	for pid := range parents {
		if err := dfs(rbac, pid, skipped, stack); err != nil {
			return err
		}
	}
	return nil
}
